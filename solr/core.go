package solr

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/at15/go-solr/solr/internal"
	"github.com/pkg/errors"
)

const (
	DefaultConfigSet = "data_driven_schema_configs"
	coreBaseURLTmpl  = "/solr/%s/"
)

type Core struct {
	Name        string `json:"name"`
	InstanceDir string `json:"instanceDir"`
	ConfigSet   string `json:"configSet"`
}

func NewCore(name string) Core {
	// TODO: actually we should valid the core name on the client side, otherwise after the creation failed, solr admin would still have a error banner
	return Core{
		Name:        name,
		InstanceDir: name,
		ConfigSet:   DefaultConfigSet,
	}
}

type CoreClient struct {
	client        *internal.Client
	solr          *Client
	name          string
	core          Core
	baseURL       string
	baseAdminURL  string
	baseUpdateURL string
	baseSelectURL string
}

func NewCoreClient(client *Client, core Core) *CoreClient {
	c := &CoreClient{
		client: client.client,
		solr:   client,
	}
	c.setCore(core)
	return c
}

func (c *CoreClient) Name() string {
	return c.name
}

func (c *CoreClient) setCore(core Core) {
	c.core = core
	c.name = core.Name
	c.baseURL = fmt.Sprintf(coreBaseURLTmpl, c.name)
	c.baseAdminURL = c.baseURL + "admin/"
	c.baseUpdateURL = c.baseURL + "update"
	c.baseSelectURL = c.baseURL + "select"
}

// http://localhost:8983/solr/demo/admin/ping?wt=json
// https://lucene.apache.org/solr/guide/6_6/ping.html
func (c *CoreClient) Ping(ctx context.Context) (*time.Duration, error) {
	start := time.Now()
	// NOTE: although we have CorePingRespone, but the QTime would be 0 by default, the QTime is the time spent on actual query
	// excluding the time for data transmission and decoding/encoding, and the default config won't do operation on real index, so
	// we count the time of the request on our client side as ping latency, which is reasonable considering the ping command we all use
	if _, err := c.client.Get(ctx, c.baseAdminURL+"ping", ioutil.Discard); err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("solr: can't ping core %s", c.name))
	}
	d := time.Now().Sub(start)
	return &d, nil
}

// http://localhost:8983/solr/admin/cores?action=STATUS&core=core-name
func (c *CoreClient) Status(ctx context.Context, includeIndexInfo bool) (*CoreStatus, error) {
	statuses, err := c.solr.CoresStatus(ctx, includeIndexInfo, c.name)
	if err != nil {
		return nil, err
	}
	status, ok := statuses[c.name]
	if !ok {
		return nil, errors.Errorf("solr: core %s not found", c.name)
	}
	return &status, nil
}
