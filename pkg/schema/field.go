package schema

const (
	FIELD_TYPE_TEXT_GENERAL = "text_general"
	FIELD_TYPE_TDATE        = "tdate" // TODO: tdate? this from the movie example
)

type Field struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	MultiValued bool   `json:"multiValued,omitempty"`
	Indexed     bool   `json:"indexed,omitempty"`
	Stored      bool   `json:"stored,omitempty"`
	Required    bool   `json:"required,omitempty"`
}
