package admin

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/internal"
	"github.com/pkg/errors"
)

const (
	coreBaseURL        = "/solr/admin/cores"
	pAction            = "action"
	pInstanceDir       = "instanceDir"
	pConfigSet         = "configSet"
	pCore              = "core"
	pName              = "name"
	pIndexInfo         = "indexInfo"
	pDeleteInstanceDir = "deleteInstanceDir"
	actionCreate       = "CREATE"
	actionStatus       = "STATUS"
	actionUnload       = "UNLOAD"
)

// http://localhost:8983/solr/admin/cores?action=CREATE&name=films&instanceDir=films&configSet=data_driven_schema_configs
func (svc *Service) CreateCore(ctx context.Context, core common.Core) error {
	// NOTE: GET works, though it should be POST
	req, err := svc.client.NewRequest(http.MethodGet, coreBaseURL, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Set(pAction, actionCreate)
	q.Set(pName, core.Name)
	if core.InstanceDir != "" {
		q.Set(pInstanceDir, core.InstanceDir)
	}
	if core.ConfigSet != "" {
		q.Set(pConfigSet, core.ConfigSet)
	}
	req.URL.RawQuery = q.Encode()
	if _, err := svc.client.Do(ctx, req, ioutil.Discard); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("solr: can't create core %s", req.URL.String()))
	}
	return nil
}

// CreateIfNotExists call Create but inspect the error message and ignore the error if the error message contains the magic words 'already exists'.
// TODO: a more robust way would be call status for specify core and then create it
func (svc *Service) CreateCoreIfNotExists(ctx context.Context, core common.Core) error {
	err := svc.CreateCore(ctx, core)
	if err != nil {
		cause := errors.Cause(err)
		if solrErr, ok := cause.(*internal.SolrErrorResponse); ok {
			if strings.Contains(solrErr.Err.Msg, "already exists") {
				return nil
			}
		}
	}
	return err
}

// http://localhost:8983/solr/admin/cores?action=STATUS
func (svc *Service) CoreStatus(ctx context.Context, indexInfo bool, core string) (map[string]*common.CoreStatus, error) {
	req, err := svc.client.NewRequest(http.MethodGet, coreBaseURL, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set(pAction, actionStatus)
	if core != "" {
		q.Set(pCore, core)
	}
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
	return res.Status, nil
}

// http://localhost:8983/solr/admin/cores?action=RENAME&core=old_name&other=new
func (svc *Service) RenameCore() error {
	return nil
}

// http://localhost:8983/solr/admin/cores?action=UNLOAD&core=my_solr_core
func (svc *Service) UnloadCore() error {
	return nil
}

// http://www.ryanwright.me/cookbook/apachesolr/delete-core
// http:/localhost:8983/solr/admin/cores?acton=UNLOAD&core=mysolrcore&deleteInstanceDir=true
func (svc *Service) DeleteCore(ctx context.Context, core string) error {
	req, err := svc.client.NewRequest(http.MethodGet, coreBaseURL, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Set(pAction, actionUnload)
	q.Set(pCore, core)
	q.Set(pDeleteInstanceDir, "true")
	req.URL.RawQuery = q.Encode()
	if _, err := svc.client.Do(ctx, req, ioutil.Discard); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("solr: can't delete core %s", req.URL.String()))
	}
	return nil
}
