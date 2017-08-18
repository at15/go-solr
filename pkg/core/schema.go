package core

import (
	"context"
	"github.com/at15/go-solr/pkg/common"
	"github.com/pkg/errors"
)

func (svc *Service) Schema(ctx context.Context) (*common.Schema, error) {
	res := &common.SchemaResponse{}
	if _, err := svc.client.Get(ctx, svc.baseURL+"schema", res); err != nil {
		return nil, errors.WithMessage(err, "solr: can't get core schmea")
	}
	return res.Schema, nil
}
