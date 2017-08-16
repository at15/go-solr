package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
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
	return nil, nil
}

// TODO: check it using http://localhost:8983/solr/admin/info/system?_=1502864003037&wt=json
// ping can only be used when a core is created https://stackoverflow.com/questions/19248746/configure-health-check-in-solr-4
func (c *Client) UseCore(core string) error {
	// TODO: there must be someway to test if a core exists or not
	return nil
}
