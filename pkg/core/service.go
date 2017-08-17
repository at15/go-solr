package core

import (
	"context"
	"fmt"
	"github.com/at15/go-solr/pkg/internal"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baseURL      = "/solr/admin/cores"
	pAction      = "action"
	pInstanceDir = "instanceDir"
	pConfigSet   = "configSet"
	pCore        = "core"
	pName        = "name"
	pIndexInfo   = "indexInfo"
	actionCreate = "CREATE"
	actionStatus = "STATUS"
)

type Service struct {
	client *internal.Client

	core Core
}

func New(client *internal.Client, core Core) *Service {
	return &Service{
		client: client,
		core:   core,
	}
}

func (svc *Service) Ping() error {
	return nil
}

// http://localhost:8983/solr/admin/cores?action=CREATE&name=films&instanceDir=films&configSet=data_driven_schema_configs
func (svc *Service) Create(ctx context.Context, core Core) error {
	req, err := svc.client.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Set(pAction, actionCreate)
	q.Set(pName, core.Name)
	q.Set(pInstanceDir, core.InstanceDir)
	q.Set(pConfigSet, core.ConfigSet)
	req.URL.RawQuery = q.Encode()
	if _, err := svc.client.Do(ctx, req, ioutil.Discard); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("solr: can't create core %s", req.URL.String()))
	}
	return nil
}

// CreateIfNotExists call Create but inspect the error message and ignore the error if the error message contains the magic words 'already exists'.
// TODO: a more robust way would be call status for specify core and then create it
func (svc *Service) CreateIfNotExists(ctx context.Context, core Core) error {
	err := svc.Create(ctx, core)
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

// http://localhost:8983/solr/admin/cores?action=STATUS&core=core-name
// TODO: extra call for getting status of all cores
func (svc *Service) Status(ctx context.Context, indexInfo bool) (*Status, error) {
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
	res := &StatusResponse{}
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

// http:localhost:PORT/solr/admin/cores?action=RENAME&core=old_name&other=new
func (svc *Service) Rename() error {
	return nil
}

// http://localhost:8983/solr/admin/cores?action=UNLOAD&core=my_solr_core
func (svc *Service) Unload() error {
	return nil
}

// http://www.ryanwright.me/cookbook/apachesolr/delete-core
// http:/localhost:8983/solr/admin/cores?acton=UNLOAD&core=mysolrcore&deleteInstanceDir=true
func (svc *Service) Delete() error {
	return nil
}
