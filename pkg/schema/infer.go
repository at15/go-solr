package schema

import (
	"reflect"
	"strings"
	"time"

	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/common/fieldtype"
	"github.com/pkg/errors"
)

const (
	TagName     = "solr"
	jsonTagName = "json"
)

var (
	typeOfByteSlice = reflect.TypeOf([]byte{})
	typeOfTime      = reflect.TypeOf(time.Time{})
)

func MustInferSchema(st interface{}) *common.Schema {
	s, err := InferSchema(st)
	if err != nil {
		log.Panic(err)
	}
	return s
}

// InferSchema generates schema based on struct definition and field tag, only Schema.Fields is generated for managed schema
// https://github.com/at15/go-solr/issues/11
func InferSchema(st interface{}) (*common.Schema, error) {
	t := reflect.TypeOf(st)
	// pointer to struct is also acceptable
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, errors.Errorf("can only infer schema from struct or pointer to struct, got %s instead", t.Kind())
	}
	sma := &common.Schema{}
	// loop through fields and ignore those unexported
	hasExportedField := false
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			log.Tracef("ignore unexported field %s", field.Name)
			continue
		}
		// ignore the field if it is ignored by either json tag or own tag
		// TODO: we need to set type for filed that is not ignored by json by ignored by solr to ignored,
		// otherwise solr seems to be automatically create field based on input document in managed schema
		if field.Tag.Get(jsonTagName) == "-" || field.Tag.Get(TagName) == "-" {
			log.Tracef("ignore field %s because - is in tag", field.Name)
			continue
		}
		hasExportedField = true
		f, err := inferFieldSchema(field)
		if err != nil {
			return nil, err
		}
		sma.Fields = append(sma.Fields, *f)
	}
	if !hasExportedField {
		return nil, errors.Errorf("type %s.%s has no exported filed", t.PkgPath(), t.Name())
	}
	return sma, nil
}

func inferFieldSchema(field reflect.StructField) (*common.Field, error) {
	//log.Tracef("infer field %s", field.Name)
	// TODO: do we need to support pointer
	// TODO: should use json name if provided
	fs := &common.Field{Name: field.Name, Type: ""}
	if err := applyTags(fs, field.Tag); err != nil {
		return nil, err
	}
	// only infer the type when user didn't explicit specify it
	if fs.Type == "" {
		switch field.Type {
		case typeOfTime:
			fs.Type = fieldtype.Date // TODO: default to trie date?
		case typeOfByteSlice:
			fs.Type = fieldtype.TextGeneral // TODO: we should document that we treat bytes slice as text general instead of binary by default
		}
		// TODO: handle slice (array)
		switch field.Type.Kind() {
		case reflect.Struct:
			if fs.Type == "" {
				return nil, errors.Errorf("nested document is not supported by go-solr, field %s is struct", field.Name)
			}
			// user provided type in tag, or it's builtin type like time handled in previous switch statement
		case reflect.Bool:
			fs.Type = fieldtype.Boolean
		case reflect.String:
			fs.Type = fieldtype.TextGeneral
		}
	}
	// no match
	if fs.Type == "" {
		return nil, errors.Errorf("unsupported field type %s", field.Type.Kind())
	}
	return fs, nil
}

func applyTags(f *common.Field, tag reflect.StructTag) error {
	//log.Tracef("all tags %v", tag)
	ApplyJSONTag(f, tag.Get(jsonTagName))
	// our own tag overrides json tag
	return ApplyTag(f, tag.Get(TagName))
}

/*
ApplyTag uses solr tag to set Name, Type and all the other attributes of a Field

	Foo	string `json:"foo" solr:",type=string,docValues=true,indexed=false,stored=true,multiValued=false,required=true"`
*/
func ApplyTag(f *common.Field, tag string) error {
	//log.Tracef("tag value %s", tag)
	values := strings.Split(tag, ",")
	if len(values) == 0 {
		return nil
	}
	if values[0] != "" && values[0] != "-" {
		f.Name = values[0]
	}
	for i := 1; i < len(values); i++ {
		kv := strings.Split(values[i], "=")
		if len(kv) != 2 {
			return errors.Errorf("invalid tag %s", values[i])
		}
		k := kv[0]
		v := kv[1]
		switch strings.ToLower(k) {
		case "type":
			// TODO: valid if this is a valid type in managed schema
			f.Type = kv[1]
		case "docvalues":
			f.DocValues = str2bool(v)
		case "indexed":
			f.Indexed = str2bool(v)
		case "stored":
			f.Stored = str2bool(v)
		case "multivalued":
			f.MultiValued = str2bool(v)
		case "required":
			f.Required = str2bool(v)
		default:
			return errors.Errorf("unknown key %s in tag %s", k, kv)
		}
	}
	return nil
}

func str2bool(s string) *bool {
	b := false
	if strings.HasPrefix(s, "t") {
		b = true
	}
	return &b
}

/*
ApplyJSONTag uses json tag to set Name, following are list of json tag values we supported base on https://godoc.org/encoding/json
	Field int `json:"myName"`
	Field int `json:"myName,omitempty"`
However, we don't support the following:
	Field int `json:"-"`
	Field int `json:"-,"`
	Int64String int64 `json:",string"`
TODO: set required = false if omitempty
*/
func ApplyJSONTag(f *common.Field, tag string) {
	//log.Tracef("json tag %s", tag)
	values := strings.Split(tag, ",")
	//log.Tracef("len is %d", len(values))
	if len(values) == 0 {
		return
	}
	if values[0] != "" && values[0] != "-" {
		f.Name = values[0]
	}
}
