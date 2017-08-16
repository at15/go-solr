package internal

import (
	"github.com/at15/go-solr/pkg/util"
)

var log = util.Logger.RegisterPkg()

const (
	DefaultUserAgent  = "go-solr"
	HeaderAccept      = "Accept"
	HeaderContentType = "Content-Type"
	MediaTypeJSON     = "application/json"
)

type Service struct {
	client *Client
}

func (svc *Service) SetClient(client *Client) {
	svc.client = client
}
