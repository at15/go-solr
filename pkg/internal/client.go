package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"fmt"
	"github.com/pkg/errors"
)

type Client struct {
	baseURL   *url.URL
	http      *http.Client
	userAgent string
}

type Response struct {
	*http.Response
}

// TODO: raw response body etc.
type SolrError struct {
	code     int
	response *SolrErrorResponse
}

type SolrErrorResponse struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
	} `json:"responseHeader"`
	// NOTE: we can't use Error because we need the method with same name to met error interface
	Err struct {
		Metadata []string `json:"metadata"`
		Msg      string   `json:"msg"`
		Trace    string   `json:"trace"`
		Code     int      `json:"code"`
	} `json:"error"`
}

func (e *SolrErrorResponse) Error() string {
	return fmt.Sprintf("solr: %d: %s", e.ResponseHeader.Status, e.Err.Msg)
}

type ClientOption func(*Client) error

func BaseURL(addr string) ClientOption {
	return func(c *Client) error {
		var err error
		c.baseURL, err = url.Parse(addr)
		if err != nil {
			return errors.Wrapf(err, "invalid base url %s", addr)
		}
		return nil
	}
}

func NewClient(client *http.Client, options ...ClientOption) (*Client, error) {
	if client == nil {
		client = http.DefaultClient
	}
	c := &Client{
		http:      client,
		userAgent: DefaultUserAgent,
	}
	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, errors.WithMessage(err, "invalid client option")
		}
	}
	return c, nil
}

// NewRequest creates a request, resolves relative url to base url, encodes body as json, sets proper request header and parameters (wt=json)
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, errors.Wrap(err, "invalid relative url")
	}
	u := c.baseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, errors.Wrap(err, "can't encode body to json")
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't create http.Request")
	}
	req.Header.Set(HeaderAccept, MediaTypeJSON)
	if body != nil {
		req.Header.Set(HeaderContentType, MediaTypeJSON)
	}
	// NOTE: Solr uses wt=json and don't check Accept header
	q := req.URL.Query()
	q.Set("wt", "json")
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func (c *Client) Get(ctx context.Context, url string, v interface{}) (*Response, error) {
	req, err := c.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(ctx, req, v)
}

func (c *Client) Post(ctx context.Context, url string, body, v interface{}) (*Response, error) {
	req, err := c.NewRequest(http.MethodGet, url, body)
	if err != nil {
		return nil, err
	}
	return c.Do(ctx, req, v)
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)
	res, err := c.http.Do(req)
	if err != nil {
		// return context error
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "context canceled")
		default:
		}
		return nil, errors.Wrap(err, "http client error")
	}
	// TODO: is this still needed in newer version of golang,
	// json decoder does not drain the body so tls connection can not be reused
	// originally from https://github.com/google/go-github/pull/317
	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, res.Body, 512)
		res.Body.Close()
	}()

	// all the non 2XX response are treated as error
	// TODO: will there be 304?
	if err = checkResponse(res); err != nil {
		return newResponse(res), err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, res.Body)
		} else {
			if err := json.NewDecoder(res.Body).Decode(v); err != nil {
				return nil, errors.Wrap(err, "can't decode json response")
			}

		}
	}
	// TODO: so what do we do when v is nil? let user drain the response body by themselves?
	return newResponse(res), nil
}

func newResponse(res *http.Response) *Response {
	return &Response{res}
}

func checkResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	// TODO: handle 404, solr return html page instead of json for route without match
	// TODO: solr seems to have a common error response
	errResp := &SolrErrorResponse{}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "can't read body")
	}
	if err := json.Unmarshal(b, errResp); err != nil {
		return errors.Wrap(err, "can't unmarshal to solr error message")
	}
	return errResp
}
