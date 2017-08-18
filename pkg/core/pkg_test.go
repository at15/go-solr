package core

import (
	"os"
	"testing"

	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/internal"
)

var tSvc *Service

func TestMain(m *testing.M) {
	log.Info("Setup of core package test")
	c := internal.MustNewInternalClient()
	tSvc = New(c, common.NewCore("demo"))
	// TODO: create demo core
	v := m.Run()
	log.Info("Tear down of core package test")
	os.Exit(v)
}
