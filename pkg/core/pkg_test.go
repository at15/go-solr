package core

import (
	"context"
	"os"
	"testing"

	"github.com/at15/go-solr/pkg/admin"
	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/internal"
)

var tSvc *Service

func TestMain(m *testing.M) {
	log.Info("Setup of core package test")
	c := internal.MustNewInternalClient()
	a := admin.New(c)
	tSvc = New(c, common.NewCore("demo"), a)
	if err := a.CreateCoreIfNotExists(context.Background(), tSvc.core); err != nil {
		log.Errorf("can't create core %v", err)
		os.Exit(1)
		return
	}
	v := m.Run()
	log.Info("Tear down of core package test")
	os.Exit(v)
}
