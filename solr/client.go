package solr

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/at15/go-solr/solr/internal"
	"github.com/pkg/errors"
)

const (
	AddrEnvName = "GO_SOLR_ADDR"
	DefaultAddr = "http://localhost:8983/"
	DefaultCore = "demo"
)

// Config
type Config struct {
	Addr        string `json:"addr" yaml:"addr"`
	DefaultCore string `json:"defaultCore" yaml:"defaultCore"`
	Cloud       bool   `json:"cloud" yaml:"cloud"`
}

// Validate checks if Addr is a valid URL and assign default values
func (c *Config) Validate() error {
	var err error
	// valid addr
	if c.Addr == "" {
		c.Addr = DefaultAddr
	}
	// addr will be used as baseURL, so it always contains a trailing slash
	if !strings.HasSuffix(c.Addr, "/") {
		c.Addr += "/"
	}
	if _, err = url.Parse(c.Addr); err != nil {
		return errors.Wrap(err, "invalid host address in config")
	}
	if c.DefaultCore == "" {
		c.DefaultCore = DefaultCore
	}
	if c.Cloud == true {
		return errors.New("SolrCloud is not supported for now")
	}
	return nil
}

// Client is used for admin tasks like create core, for document operation (core specific) like indexing, query use CoreClient instead
type Client struct {
	mu     sync.Mutex
	config Config
	client *internal.Client

	DefaultCore *CoreClient
	cores       map[string]*CoreClient
}

// NewClient creates Client based on config, it returns error when the config is invalid, it does NOT check network connection
func NewClient(config Config) (*Client, error) {
	var err error
	if err = config.Validate(); err != nil {
		return nil, err
	}
	c := &Client{
		config: config,
		cores:  make(map[string]*CoreClient),
	}
	// TODO: let user config transport (i.e. use ss) and client timeout
	tr := &http.Transport{}
	h := &http.Client{Transport: tr}
	if c.client, err = internal.NewClient(h, internal.BaseURL(config.Addr)); err != nil {
		return nil, errors.WithMessage(err, "can't create internal http client wrapper")
	}
	c.DefaultCore = NewCoreClient(c, NewCore(config.DefaultCore))
	c.cores[config.DefaultCore] = c.DefaultCore
	return c, nil
}

// IsUp check if the server is running by querying system info, we don't use ping because you can only ping a core, not entire solr
// https://stackoverflow.com/questions/19248746/configure-health-check-in-solr-4
func (c *Client) IsUp(ctx context.Context) error {
	// using http://localhost:8983/solr/admin/info/system?wt=json
	info, err := c.SystemInfo(ctx)
	log.Debug(info)
	return err
}

// UseCore change DefaultCore to CoreClient with coreName, it does NOT check if the core exists or is up
func (c *Client) UseCore(coreName string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.DefaultCore = NewCoreClient(c, NewCore(coreName))
	c.cores[coreName] = c.DefaultCore
}

// GetCore returns existing CoreClient or create a new one and save it before return
func (c *Client) GetCore(coreName string) *CoreClient {
	c.mu.Lock()
	defer c.mu.Unlock()

	cor, ok := c.cores[coreName]
	if ok {
		return cor
	}
	cor = NewCoreClient(c, NewCore(coreName))
	c.cores[coreName] = cor
	return cor
}
