package admin

import (
	"context"
	"github.com/at15/go-solr/pkg/internal"
)

type Service struct {
	client *internal.Client
}

func New(client *internal.Client) *Service {
	return &Service{
		client: client,
	}
}

func (svc *Service) SystemInfo(ctx context.Context) (*SystemInfoResponse, error) {
	res := &SystemInfoResponse{}
	if _, err := svc.client.Get(ctx, baseURL+"/info/system", res); err != nil {
		return res, err
	}
	return res, nil
}
