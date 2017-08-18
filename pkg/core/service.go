package core

import (
	"context"
	"fmt"
	"net/http"

	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/internal"
	"github.com/pkg/errors"
)

const (
	baseURL      = "/solr/admin/cores"
	pAction      = "action"
	pCore        = "core"
	pIndexInfo   = "indexInfo"
	actionStatus = "STATUS"
)

type Service struct {
	client *internal.Client

	core common.Core
}

func New(client *internal.Client, core common.Core) *Service {
	return &Service{
		client: client,
		core:   core,
	}
}

func (svc *Service) Ping() error {
	return nil
}

// http://localhost:8983/solr/admin/cores?action=STATUS&core=core-name
// TODO: extra call for getting status of all cores
// TODO: move this to admin/core.go
func (svc *Service) Status(ctx context.Context, indexInfo bool) (*common.CoreStatus, error) {
	req, err := svc.client.NewRequest(http.MethodGet, baseURL, nil)
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
