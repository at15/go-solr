package core

import (
	"context"

	"github.com/at15/go-solr/pkg/common"
	"github.com/pkg/errors"
)

// http://localhost:8983/solr/demo/admin/ping?wt=json
// https://lucene.apache.org/solr/guide/6_6/ping.html
func (svc *Service) Ping(ctx context.Context) error {
	res := &common.CorePingResponse{}
	if _, err := svc.client.Get(ctx, svc.baseAdminURL+"ping", res); err != nil {
		// TODO: we want to log the request url, maybe this should be included in client's error message generation
		return errors.WithMessage(err, "solr: can't ping core %s")
	}
	return nil
}

// http://localhost:8983/solr/admin/cores?action=STATUS&core=core-name
func (svc *Service) Status(ctx context.Context, indexInfo bool) (*common.CoreStatus, error) {
	allStatus, err := svc.admin.CoreStatus(ctx, indexInfo, svc.core.Name)
	if err != nil {
		return nil, err
	}
	status, ok := allStatus[svc.core.Name]
	if !ok {
		return nil, errors.Errorf("solr: core %s not found %s", svc.core.Name)
	}
	return status, nil
}
