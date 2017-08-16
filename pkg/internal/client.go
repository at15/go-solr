package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"io/ioutil"
)

type Client struct {
	baseURL   url.URL
	http      *http.Client
	userAgent string
}

type Response struct {
	*http.Response
}

func NewClient(c *http.Client) *Client {
	if c == nil {
		c = http.DefaultClient
	}
	return &Client{
		http:      c,
		userAgent: DefaultUserAgent,
	}
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
	return nil, nil
}

func (c *Client) Get(ctx context.Context, url string, v interface{}) (*Response, error) {
	req, err := c.NewRequest(http.MethodGet, url, nil)
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
	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, res.Body, 512)
		res.Body.Close()
	}()
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, res.Body)
		} else {
			if err := json.NewDecoder(res.Body).Decode(v); err != nil {
				return nil, errors.Wrap(err, "can't decode json response")
			}

		}
	}
	return newResponse(res), nil
}

func newResponse(res *http.Response) *Response {
	return &Response{res}
}

// TODO: check it using http://localhost:8983/solr/admin/info/system?_=1502864003037&wt=json
// ping can only be used when a core is created https://stackoverflow.com/questions/19248746/configure-health-check-in-solr-4
func (c *Client) UseCore(core string) error {
	// TODO: there must be someway to test if a core exists or not
	return nil
}
