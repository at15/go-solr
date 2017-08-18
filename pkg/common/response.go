package common

import (
	"time"
)

// They are semi auto generated from json using https://mholt.github.io/json-to-go/

// CoreStatusResponse is core status (with index) returned from http://localhost:8983/solr/admin/cores?action=STATUS&wt=json&core=demo
type CoreStatusResponse struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
	} `json:"responseHeader"`
	InitFailures interface{}            `json:"initFailures"`
	Status       map[string]*CoreStatus `json:"status"`
}

// CoreStatus is core status returned from http://localhost:8983/solr/admin/cores?action=STATUS&wt=json&core=demo
type CoreStatus struct {
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

// CorePingResponse is returned from http://localhost:8983/solr/demo/admin/ping?wt=json
type CorePingResponse struct {
	ResponseHeader struct {
		ZkConnected interface{} `json:"zkConnected"`
		Status      int         `json:"status"`
		QTime       int         `json:"QTime"`
		Params      struct {
			Q          string `json:"q"`
			Distrib    string `json:"distrib"`
			Df         string `json:"df"`
			Rows       string `json:"rows"`
			Wt         string `json:"wt"`
			EchoParams string `json:"echoParams"`
		} `json:"params"`
	} `json:"responseHeader"`
	Status string `json:"status"`
}

// SystemInfoResponse is system info returned from http://localhost:8983/solr/admin/info/system?wt=json
// NOTE: we use interface{} for systemCpuLoad because it is "NaN" on Mac, and can't be decoded as float64
type SystemInfoResponse struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
	} `json:"responseHeader"`
	Mode     string `json:"mode"`
	SolrHome string `json:"solr_home"`
	Lucene   struct {
		SolrSpecVersion   string `json:"solr-spec-version"`
		SolrImplVersion   string `json:"solr-impl-version"`
		LuceneSpecVersion string `json:"lucene-spec-version"`
		LuceneImplVersion string `json:"lucene-impl-version"`
	} `json:"lucene"`
	Jvm struct {
		Version string `json:"version"`
		Name    string `json:"name"`
		Spec    struct {
			Vendor  string `json:"vendor"`
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"spec"`
		Jre struct {
			Vendor  string `json:"vendor"`
			Version string `json:"version"`
		} `json:"jre"`
		VM struct {
			Vendor  string `json:"vendor"`
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"vm"`
		Processors int `json:"processors"`
		Memory     struct {
			Free  string `json:"free"`
			Total string `json:"total"`
			Max   string `json:"max"`
			Used  string `json:"used"`
			Raw   struct {
				Free        int     `json:"free"`
				Total       int     `json:"total"`
				Max         int     `json:"max"`
				Used        int     `json:"used"`
				UsedPercent float64 `json:"used%"`
			} `json:"raw"`
		} `json:"memory"`
		Jmx struct {
			Bootclasspath   string    `json:"bootclasspath"`
			Classpath       string    `json:"classpath"`
			CommandLineArgs []string  `json:"commandLineArgs"`
			StartTime       time.Time `json:"startTime"`
			UpTimeMS        int       `json:"upTimeMS"`
		} `json:"jmx"`
	} `json:"jvm"`
	System struct {
		Name                       string      `json:"name"`
		Arch                       string      `json:"arch"`
		AvailableProcessors        int         `json:"availableProcessors"`
		SystemLoadAverage          float64     `json:"systemLoadAverage"`
		Version                    string      `json:"version"`
		CommittedVirtualMemorySize int64       `json:"committedVirtualMemorySize"`
		FreePhysicalMemorySize     int64       `json:"freePhysicalMemorySize"`
		FreeSwapSpaceSize          int64       `json:"freeSwapSpaceSize"`
		ProcessCPULoad             float64     `json:"processCpuLoad"`
		ProcessCPUTime             int64       `json:"processCpuTime"`
		SystemCPULoad              interface{} `json:"systemCpuLoad"` // FIXME: on Mac this is "NaN" and golang can't decode it as float64
		TotalPhysicalMemorySize    int64       `json:"totalPhysicalMemorySize"`
		TotalSwapSpaceSize         int64       `json:"totalSwapSpaceSize"`
		MaxFileDescriptorCount     int         `json:"maxFileDescriptorCount"`
		OpenFileDescriptorCount    int         `json:"openFileDescriptorCount"`
		Uname                      string      `json:"uname"`
		Uptime                     string      `json:"uptime"`
	} `json:"system"`
}
