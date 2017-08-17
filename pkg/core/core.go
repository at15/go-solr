package core

import "time"

// https://cwiki.apache.org/confluence/display/solr/CoreAdmin+API
// NOTE: must specify configSet
// if you're running SolrCloud, you should NOT be using the CoreAdmin API at all. Use the Collections API.

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

// TODO: NewCoreFromStatus

// StatusResponse is core status (with index) returned from /solr/admin/cores?action=STATUS&wt=json&core=demo
type StatusResponse struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
	} `json:"responseHeader"`
	InitFailures interface{}        `json:"initFailures"`
	Status       map[string]*Status `json:"status"`
}

// Status is core status returned from /solr/admin/cores?action=STATUS&wt=json&core=demo
// it is generated using https://mholt.github.io/json-to-go/
type Status struct {
	Name        string    `json:"name"`
	InstanceDir string    `json:"instanceDir"`
	DataDir     string    `json:"dataDir"`
	Config      string    `json:"config"`
	Schema      string    `json:"schema"`
	StartTime   time.Time `json:"startTime"`
	Uptime      int       `json:"uptime"`
	Index       struct {
		NumDocs                 int    `json:"numDocs"`
		MaxDoc                  int    `json:"maxDoc"`
		DeletedDocs             int    `json:"deletedDocs"`
		IndexHeapUsageBytes     int    `json:"indexHeapUsageBytes"`
		Version                 int    `json:"version"`
		SegmentCount            int    `json:"segmentCount"`
		Current                 bool   `json:"current"`
		HasDeletions            bool   `json:"hasDeletions"`
		Directory               string `json:"directory"`
		SegmentsFile            string `json:"segmentsFile"`
		SegmentsFileSizeInBytes int    `json:"segmentsFileSizeInBytes"`
		UserData                struct {
		} `json:"userData"`
		SizeInBytes int    `json:"sizeInBytes"`
		Size        string `json:"size"`
	} `json:"index"`
}
