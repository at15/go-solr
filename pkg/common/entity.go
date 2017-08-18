package common

const (
	DefaultConfigSet = "data_driven_schema_configs"
)

type Core struct {
	Name        string `json:"name"`
	InstanceDir string `json:"instanceDir"`
	ConfigSet   string `json:"configSet"`
}

func NewCore(name string) Core {
	// TODO: actually we should valid the core name on the client side, otherwise after the creation failed, solr admin would still have a error banner
	return Core{
		Name:        name,
		InstanceDir: name,
		ConfigSet:   DefaultConfigSet,
	}
}
