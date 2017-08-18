package core

import (
	"context"
	"fmt"
	"net/http"

	"github.com/at15/go-solr/pkg/common"
	"github.com/pkg/errors"
)

const (
	globalAdminURL = "/solr/admin/cores"
	pAction        = "action"
	pCore          = "core"
	pIndexInfo     = "indexInfo"
	actionStatus   = "STATUS"
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
// FIXME: this is almost identical to admin/core.go, but we can't import admin service
func (svc *Service) Status(ctx context.Context, indexInfo bool) (*common.CoreStatus, error) {
	req, err := svc.client.NewRequest(http.MethodGet, globalAdminURL, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set(pAction, actionStatus)
	q.Set(pCore, svc.core.Name)
	if indexInfo {
		q.Set(pIndexInfo, "true")
	} else {
		q.Set(pIndexInfo, "false")
	}
	req.URL.RawQuery = q.Encode()
	res := &common.CoreStatusResponse{}
	if _, err := svc.client.Do(ctx, req, res); err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("solr: can't get core status %s", req.URL.String()))
	}
	// even though we specify the core, solr still wraps it with a map
	status, ok := res.Status[svc.core.Name]
	if !ok {
		return nil, errors.Errorf("solr: core %s not found %s", svc.core.Name, req.URL.String())
	}
	return status, nil
}
