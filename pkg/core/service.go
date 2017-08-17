package core

import "github.com/at15/go-solr/pkg/internal"

type Service struct {
	client *internal.Client

	core *Core
}

func New(client *internal.Client, core *Core) *Service {
	return &Service{
		client: client,
		core:   core,
	}
}

func (svc *Service) Ping() error {
	return nil
}

func (svc *Service) Create() error {
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
