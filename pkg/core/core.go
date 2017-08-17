package core

// https://cwiki.apache.org/confluence/display/solr/CoreAdmin+API
// http://localhost:8983/solr/admin/cores?action=CREATE&name=films&instanceDir=films&configSet=data_driven_schema_configs
// NOTE: must specify configSet
// if you're running SolrCloud, you should NOT be using the CoreAdmin API at all. Use the Collections API.

const (
	DefaultConfigSet = "data_driven_schema_configs"
)

// TODO: move this to the core package
type Core struct {
	Name        string `json:"name"`
	InstanceDir string `json:"instanceDir"`
	ConfigSet   string `json:"configSet"`
}

func NewCore(name string) *Core {
	// TODO: actually we should valid the core name on the client side, otherwise after the creation failed, solr admin would still have a error banner
	return &Core{
		Name:        name,
		InstanceDir: name,
		ConfigSet:   DefaultConfigSet,
	}
}
