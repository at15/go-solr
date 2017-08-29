package admin

import (
	"os"
	"testing"

	"github.com/at15/go-solr/pkg/internal"
)

var tSvc *Service

func TestMain(m *testing.M) {
	log.Info("Setup of admin package test")
	c := internal.MustNewInternalClient()
	tSvc = New(c)
	v := m.Run()
	log.Info("Tear down of admin package test")
	os.Exit(v)
}
