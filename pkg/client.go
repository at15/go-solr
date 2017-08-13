package pkg

import (
	"github.com/at15/go-solr/pkg/util"
	"net/http"
)

var log = util.Logger.RegisterPkg()

type Client struct {
	client *http.Client
}

func New(c *http.Client) *Client {
	if c == nil {
		c = http.DefaultClient
	}
	return &Client{
		client: c,
	}
}
