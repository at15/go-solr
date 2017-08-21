package common

import "encoding/json"

const (
	DefaultConfigSet = "data_driven_schema_configs"
)

var (
	// True is for setting *bool, i.e. field.docValues = &True, we use *bool instead of bool to avoid omitempty ignore false value
	True  = true
	False = false
)

type Document interface {
	json.Marshaler
}

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

type Schema struct {
	Name          string         `json:"name"`
	Version       float64        `json:"version"`
	UniqueKey     string         `json:"uniqueKey"`
	FieldTypes    []FieldType    `json:"fieldTypes"`
	Fields        []Field        `json:"fields"`
	DynamicFields []DynamicField `json:"dynamicFields"`
	CopyFields    []CopyField    `json:"copyFields"`
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
