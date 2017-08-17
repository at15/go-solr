// Package core contains operations related to Solr Core
package core

import (
	"github.com/at15/go-solr/pkg/util"
)

var log = util.Logger.RegisterPkg()

const (
	baseURL = "/solr/admin/cores"
)
