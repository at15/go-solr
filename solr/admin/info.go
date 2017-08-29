package admin

import (
	"context"
	"github.com/at15/go-solr/pkg/common"
)

func (svc *Service) SystemInfo(ctx context.Context) (*common.SystemInfoResponse, error) {
	res := &common.SystemInfoResponse{}
	if _, err := svc.client.Get(ctx, baseURL+"/info/system", res); err != nil {
		return res, err
	}
	return res, nil
}
