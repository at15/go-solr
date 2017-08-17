package internal

import (
	"github.com/at15/go-solr/pkg/util"
)

var log = util.Logger.RegisterPkg()

const (
	DefaultUserAgent  = "go-solr"
	HeaderAccept      = "Accept"
	HeaderContentType = "Content-Type"
	MediaTypeJSON     = "application/json"
)
