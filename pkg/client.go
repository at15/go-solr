package pkg

import (
	"context"

	"github.com/at15/go-solr/pkg/admin"
	"github.com/at15/go-solr/pkg/core"
	"github.com/at15/go-solr/pkg/internal"
	"github.com/at15/go-solr/pkg/schema"
	"github.com/at15/go-solr/pkg/util"
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

var log = util.Logger.RegisterPkg()

const (
	DefaultAddr = "http://localhost:8983/"
)

type Config struct {
	Addr  string `json:"addr" yaml:"addr"`
	Cloud bool   `json:"cloud" yaml:"cloud"`
}

type SolrClient struct {
	config Config
	client *internal.Client

	common internal.Service // reuse single struct

	Admin  *admin.Service
	Core   *core.Service // TODO: might have more than one core services to allow user use multiple cores simut
	Schema *schema.Service
}

func New(config Config) (*SolrClient, error) {
	// valid addr
	if config.Addr == "" {
		config.Addr = DefaultAddr
	}
	// addr will be used as baseURL, so it always contains a trailing slash
	if !strings.HasSuffix(config.Addr, "/") {
		config.Addr += "/"
	}
	if _, err := url.Parse(config.Addr); err != nil {
		return nil, errors.Wrap(err, "invalid host address in config")
	}
	c := &SolrClient{config: config}
	// TODO: our default behaviour should be create a new transport and set timeout to the http client instead of using
	// the default transport and client
	c.client = internal.NewClient(nil)
	c.Admin = admin.New(c.client)
	// TODO: remove the usage of casting from common service
	c.common.SetClient(c.client)
	c.Core = (*core.Service)(&c.common)
	c.Schema = (*schema.Service)(&c.common)
	return c, nil
}

// TODO: check it using http://localhost:8983/solr/admin/info/system?_=1502864003037&wt=json
// ping can only be used when a core is created https://stackoverflow.com/questions/19248746/configure-health-check-in-solr-4
func (c *SolrClient) IsUp(ctx context.Context) error {
	info, err := c.Admin.SystemInfo(ctx)
	log.Info(info)
	return err
}

func (c *SolrClient) UseCore(core string) error {
	// TODO: there must be someway to test if a core exists or not
	return nil
}
