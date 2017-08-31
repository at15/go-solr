package solr

import (
	"encoding/json"
	"time"
)

// Response structs are semi auto generated from json using https://mholt.github.io/json-to-go/ ,
// we need to change some fields to interface{} to make sure unmarshal won't break

// CoreStatusResponse is status of all the cores or single core (if specified) returned from http://localhost:8983/solr/admin/cores?action=STATUS&wt=json&core=demo
type CoreStatusResponse struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
	} `json:"responseHeader"`
	InitFailures interface{}           `json:"initFailures"`
	Status       map[string]CoreStatus `json:"status"`
}

// CoreStatus contains index information returned from http://localhost:8983/solr/admin/cores?action=STATUS&wt=json&core=demo
type CoreStatus struct {
	Name        string    `json:"name"`
	InstanceDir string    `json:"instanceDir"`
	DataDir     string    `json:"dataDir"`
	Config      string    `json:"config"`
	Schema      string    `json:"schema"`
	StartTime   time.Time `json:"startTime"`
	Uptime      int       `json:"uptime"`
	Index       struct {
		NumDocs                 int         `json:"numDocs"`
		MaxDoc                  int         `json:"maxDoc"`
		DeletedDocs             int         `json:"deletedDocs"`
		IndexHeapUsageBytes     int         `json:"indexHeapUsageBytes"`
		Version                 int         `json:"version"`
		SegmentCount            int         `json:"segmentCount"`
		Current                 bool        `json:"current"`
		HasDeletions            bool        `json:"hasDeletions"`
		Directory               string      `json:"directory"`
		SegmentsFile            string      `json:"segmentsFile"`
		SegmentsFileSizeInBytes int         `json:"segmentsFileSizeInBytes"`
		UserData                interface{} `json:"userData"`
		SizeInBytes             int         `json:"sizeInBytes"`
		Size                    string      `json:"size"`
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

type FieldType struct {
	Name          string `json:"name"`
	Class         string `json:"class"`
	IndexAnalyzer struct {
		Tokenizer struct {
			Class string `json:"class"`
		} `json:"tokenizer"`
	} `json:"indexAnalyzer,omitempty"`
	QueryAnalyzer struct {
		Tokenizer struct {
			Class     string `json:"class"`
			Delimiter string `json:"delimiter"`
		} `json:"tokenizer"`
	} `json:"queryAnalyzer,omitempty"`
	SortMissingLast      *bool  `json:"sortMissingLast,omitempty"`
	MultiValued          *bool  `json:"multiValued,omitempty"`
	CurrencyConfig       string `json:"currencyConfig,omitempty"`
	DefaultCurrency      string `json:"defaultCurrency,omitempty"`
	PrecisionStep        string `json:"precisionStep,omitempty"`
	PositionIncrementGap string `json:"positionIncrementGap,omitempty"`
	DocValues            *bool  `json:"docValues,omitempty"`
	Indexed              *bool  `json:"indexed,omitempty"`
	Stored               *bool  `json:"stored,omitempty"`
	Analyzer             struct {
		Tokenizer struct {
			Class string `json:"class"`
		} `json:"tokenizer"`
		Filters []struct {
			Class   string `json:"class"`
			Encoder string `json:"encoder"`
		} `json:"filters"`
	} `json:"analyzer,omitempty"`
	Geo                       string `json:"geo,omitempty"`
	MaxDistErr                string `json:"maxDistErr,omitempty"`
	DistErrPct                string `json:"distErrPct,omitempty"`
	DistanceUnits             string `json:"distanceUnits,omitempty"`
	SubFieldSuffix            string `json:"subFieldSuffix,omitempty"`
	Dimension                 string `json:"dimension,omitempty"`
	AutoGeneratePhraseQueries string `json:"autoGeneratePhraseQueries,omitempty"`
}

// Field uses pointer for bool to create proper payload, we omit not specified flags and let server handle them,
// so we are less likely to run into inconsistent default flags in go-solr and solr itself
type Field struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	DocValues   *bool  `json:"docValues,omitempty"`
	Indexed     *bool  `json:"indexed,omitempty"`
	Stored      *bool  `json:"stored,omitempty"`
	MultiValued *bool  `json:"multiValued,omitempty"`
	Required    *bool  `json:"required,omitempty"`
}

func NewField(name string, fieldType string) Field {
	return Field{
		Name: name,
		Type: fieldType,
	}
}

type DynamicField struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Indexed     *bool  `json:"indexed,omitempty"`
	Stored      *bool  `json:"stored,omitempty"`
	MultiValued *bool  `json:"multiValued,omitempty"`
}

type CopyField struct {
	Source string `json:"source"`
	Dest   string `json:"dest"`
}

type Schema struct {
	Name          string         `json:"name"`
	Version       float64        `json:"version"`
	UniqueKey     string         `json:"uniqueKey"`
	FieldTypes    []FieldType    `json:"fieldTypes"`
	Fields        []Field        `json:"fields"`
	DynamicFields []DynamicField `json:"dynamicFields"`
	CopyFields    []CopyField    `json:"copyFields"`
}

// SchemaResponse is schema information of a single core returned from http://localhost:8983/solr/demo/schema?wt=json
type SchemaResponse struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
	} `json:"responseHeader"`
	Schema *Schema `json:"schema"`
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

type SelectResponse struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
		Params struct {
			Q          string      `json:"q"`
			Df         string      `json:"df"`
			FacetField interface{} `json:"facet.field"` // NOTE: this could be string/[]string
			Facet      string      `json:"facet"`
			Indent     string      `json:"indent"`
			Start      string      `json:"start"`
			Sort       string      `json:"sort"`
			Wt         string      `json:"wt"`
			//NAMING_FAILED string `json:"_"`
		} `json:"params"`
	} `json:"responseHeader"`
	Response struct {
		NumFound int                      `json:"numFound"`
		Start    int                      `json:"start"`
		Docs     []map[string]interface{} `json:"docs"`
	} `json:"response"`
	FacetCounts struct {
		FacetQueries   interface{}           `json:"facet_queries"`
		FacetFields    map[string]FacetField `json:"facet_fields"` // NOTE: facet fields mix string and number in array https://github.com/at15/go-solr/issues/17
		FacetRanges    interface{}           `json:"facet_ranges"`
		FacetIntervals interface{}           `json:"facet_intervals"`
		FacetHeatmaps  interface{}           `json:"facet_heatmaps"`
	} `json:"facet_counts"`
}

type FacetField struct {
	Values []string `json:"values"`
	Counts []int    `json:"counts"`
}

// https://groups.google.com/forum/#!topic/golang-nuts/IxPipKwI-zQ
// https://play.golang.org/p/YgUIFxT7hA
func (f *FacetField) UnmarshalJSON(data []byte) error {
	//log.Info("facet field json unmarshaler called")
	var mixed []json.Number
	if err := json.Unmarshal(data, &mixed); err != nil {
		return err
	}
	f.Values = make([]string, 0, len(mixed)/2)
	f.Counts = make([]int, 0, len(mixed)/2)
	for i := 0; i < len(mixed); i += 2 {
		f.Values = append(f.Values, mixed[i].String())
		c, err := mixed[i+1].Int64()
		if err != nil {
			return err
		}
		f.Counts = append(f.Counts, int(c))
	}
	return nil
}
