package solr

import (
	"context"

	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/at15/go-solr/pkg/admin"
	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/core"
	"github.com/at15/go-solr/solr/internal"
	"github.com/at15/go-solr/solr/util/logutil"
	"github.com/pkg/errors"
)

var log = logutil.Logger.RegisterPkg()

const (
	DefaultAddr = "http://localhost:8983/"
	DefaultCore = "demo"
)

type Config struct {
	Addr        string `json:"addr" yaml:"addr"`
	DefaultCore string `json:"defaultCore" yaml:"defaultCore"`
	Cloud       bool   `json:"cloud" yaml:"cloud"`
}

func (c *Confg) Validate() error {
	// valid addr
	if config.Addr == "" {
		config.Addr = DefaultAddr
	}
	// addr will be used as baseURL, so it always contains a trailing slash
	if !strings.HasSuffix(config.Addr, "/") {
		config.Addr += "/"
	}
	if _, err = url.Parse(config.Addr); err != nil {
		return nil, errors.Wrap(err, "invalid host address in config")
	}
	if config.DefaultCore == "" {
		config.DefaultCore = DefaultCore
	}
}

type SolrClient struct {
	mu     sync.Mutex
	config Config
	client *internal.Client

	DefaultCore *core.Service
	cores       map[string]*core.Service
}

func NewClient(config Config) (*SolrClient, error) {
	var err error
	if err = config.Validate(); err != nil {
		return nil, err
	}
	c := &SolrClient{
		config: config,
		cores:  make(map[string]*core.Service),
	}
	// TODO: let user config transport (i.e. use ss) and client timeout
	tr := &http.Transport{}
	h := &http.Client{Transport: tr}
	if c.client, err = internal.NewClient(h, internal.BaseURL(config.Addr)); err != nil {
		return nil, errors.WithMessage(err, "can't create internal http client wrapper")
	}
	c.Admin = admin.New(c.client)
	c.DefaultCore = core.New(c.client, common.NewCore(config.DefaultCore), c.Admin)
	c.cores[config.DefaultCore] = c.DefaultCore
	return c, nil
}

// ping can only be used when a core is created https://stackoverflow.com/questions/19248746/configure-health-check-in-solr-4
func (c *SolrClient) IsUp(ctx context.Context) error {
	// using http://localhost:8983/solr/admin/info/system?wt=json
	info, err := c.Admin.SystemInfo(ctx)
	log.Debug(info)
	return err
}

func (c *SolrClient) UseCore(coreName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.DefaultCore = core.New(c.client, common.NewCore(coreName), c.Admin)
	c.cores[c.DefaultCore.NameOfCore()] = c.DefaultCore

	// TODO: maybe we should test if this core exists
	return nil
}

func (c *SolrClient) GetCore(coreName string) (*core.Service, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cor, ok := c.cores[coreName]
	if ok {
		return cor, nil
	}
	cor = core.New(c.client, common.NewCore(coreName), c.Admin)
	c.cores[coreName] = cor
	return cor, nil
}
