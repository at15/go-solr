// Package core contains operations related to Solr Core.
// Official API doc is http://lucene.apache.org/solr/guide/6_6/coreadmin-api.html
package core

import (
	"github.com/at15/go-solr/pkg/util"
)

var log = util.Logger.RegisterPkg()

// - [ ] TODO: solrj pass core in url when create client
