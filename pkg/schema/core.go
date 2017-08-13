package schema

// https://cwiki.apache.org/confluence/display/solr/CoreAdmin+API
// http://localhost:8983/solr/admin/cores?action=CREATE&name=films&instanceDir=films&configSet=data_driven_schema_configs
// NOTE: must specify configSet
// if you're running SolrCloud, you should NOT be using the CoreAdmin API at all. Use the Collections API.

const (
	DEAFULT_CONFIGSET = "data_driven_schema_configs"
)
type CreateCoreRequest struct {
	Name        string `json:"name"`
	InstanceDir string `json:"instanceDir"`
	ConfigSet   string `json:"configSet"`
}
