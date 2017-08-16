package admin

import (
	"context"
	"github.com/at15/go-solr/pkg/internal"
)

type Service internal.Service

func (svc *Service) SystemInfo(ctx context.Context) (*SystemInfoResponse, error) {
	res := &SystemInfoResponse{}
	if _, err := svc.client.Get(ctx, baseURL + "/info/system", res); err != nil {
		return res, err
	}
	return res, nil
}
