package admin

import "time"

// SystemInfoResponse is returned from /solr/admin/info/system?wt=json
// it is semi auto generated from json using https://mholt.github.io/json-to-go/
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
