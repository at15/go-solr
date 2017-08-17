package schema

import "github.com/at15/go-solr/pkg/internal"

type Service struct {
	client *internal.Client
}

func New(client *internal.Client) *Service {
	return &Service{
		client: client,
	}
}
