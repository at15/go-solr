package core

import (
	"context"
	"fmt"
	"github.com/at15/go-solr/pkg/internal"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const (
	baseURL      = "/solr/admin/cores"
	pAction      = "action"
	pInstanceDir = "instanceDir"
	pConfigSet   = "configSet"
	pCore        = "core"
	pName        = "name"
	actionCreate = "CREATE"
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
		return errors.WithMessage(err, fmt.Sprintf("can't create core %s", req.URL.String()))
	}
	return nil
}

// http://localhost:8983/solr/admin/cores?action=STATUS&core=core-name
// TODO: extra call for getting status of all cores
// TODO: core status struct
func (svc *Service) Status() error {
	return nil
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
