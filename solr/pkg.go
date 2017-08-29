package solr

import (
	"github.com/at15/go-solr/solr/util/logutil"
)

var log = logutil.Logger.RegisterPkg()

func init() {
	log.SetPkgAlias("solr")
}
