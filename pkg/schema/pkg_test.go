package schema

import (
	"context"
	"os"
	"testing"

	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/internal"
	asst "github.com/stretchr/testify/assert"
)

var tSvc *Service

func TestMain(m *testing.M) {
	log.Info("Setup of schema package test")
	c := internal.MustNewInternalClient()
	tSvc = New(c, common.NewCore("demo"))
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
