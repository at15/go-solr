// Package core contains operations related to Solr Core.
// Official API doc is http://lucene.apache.org/solr/guide/6_6/coreadmin-api.html
package core

import (
	"fmt"

	"github.com/at15/go-solr/pkg/admin"
	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/internal"
	"github.com/at15/go-solr/pkg/schema"
	"github.com/at15/go-solr/pkg/util"
)

const (
	baseURLTmpl = "/solr/%s/"
)

var log = util.Logger.RegisterPkg()

// - [ ] TODO: solrj pass core in url when create client
type Service struct {
	client *internal.Client
	admin  *admin.Service
	Schema *schema.Service

	core         common.Core
	baseURL      string
	baseAdminURL string
}

func New(client *internal.Client, core common.Core, admin *admin.Service) *Service {
	s := schema.New(client, core)
	svc := &Service{
		client: client,
		admin:  admin,
		Schema: s,
	}
	svc.setCore(core)
	return svc
}

func (svc *Service) setCore(core common.Core) {
	svc.core = core
	svc.baseURL = fmt.Sprintf(baseURLTmpl, core.Name)
	svc.baseAdminURL = svc.baseURL + "admin/"
}
