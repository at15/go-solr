package admin

import (
	"github.com/at15/go-solr/pkg/internal"
	"github.com/at15/go-solr/pkg/util"
)

var log = util.Logger.RegisterPkg()

const (
	baseURL = "/solr/admin"
)

type Service struct {
	client *internal.Client
}

func New(client *internal.Client) *Service {
	return &Service{
		client: client,
	}
}
