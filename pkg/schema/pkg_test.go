package schema

import (
	"context"
	"os"
	"testing"

	"github.com/at15/go-solr/pkg/admin"
	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/internal"
	asst "github.com/stretchr/testify/assert"
)

var tSvc *Service

func TestMain(m *testing.M) {
	log.Info("Setup of schema package test")
	c := internal.MustNewInternalClient()
	a := admin.New(c)
	tSvc = New(c, common.NewCore("demoschema"))
	if err := a.CreateCoreIfNotExists(context.Background(), tSvc.core); err != nil {
		log.Errorf("can't create core %v", err)
		os.Exit(1)
		return
	}
	v := m.Run()
	log.Info("Tear down of schema package test")
	os.Exit(v)
}

func TestService_Get(t *testing.T) {
	assert := asst.New(t)

	s, err := tSvc.Get(context.Background())
	assert.Nil(err)
	assert.Equal("example-data-driven-schema", s.Name)
	assert.Equal("id", s.UniqueKey)
}
