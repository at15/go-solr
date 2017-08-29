package fieldtype

const (
	Ignored = "ignored"
)

// primitive types
const (
	Binary       = "binary"
	Boolean      = "boolean"
	MultiBoolean = "booleans"
	// TODO: int
)

// date
const (
	// Date is solr.TrieDateField with precisionStep as 0
	Date      = "date"
	MultiDate = "dates"
	// TrieDate is solr.TrieDateField with precisionStep as 6
	TrieDate      = "tdate"
	MultiTrieDate = "tdates"
)

// text
const (
	// String is UTF-8 encoded string NOT tokenized or analyzed in any way, hard limit is 32K
	String      = "string"
	MultiString = "strings"
	// TextEn is text with english locale, its multiValued is set to false
	TextEn = "text_en"
	// TextGeneral is text without locale, TODO: its multiValued is set to true
	TextGeneral = "text_general"
)
