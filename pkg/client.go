package pkg

import (
	"github.com/at15/go-solr/pkg/util"
	"net/http"
)

var log = util.Logger.RegisterPkg()

const (
	DefaultHost = "localhost"
	DefaultPort = 8983
)
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

// TODO: check it using http://localhost:8983/solr/admin/info/system?_=1502864003037&wt=json
// ping can only be used when a core is created https://stackoverflow.com/questions/19248746/configure-health-check-in-solr-4