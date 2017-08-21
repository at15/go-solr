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
	applyTags(fs, field.Tag)
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
	if fs.Type == "" {
		return nil, errors.Errorf("unsupported field type %s", field.Type.Kind())
	}
	return fs, nil
}

func applyTags(f *common.Field, tag reflect.StructTag) {
	log.Tracef("tag value %s", tag.Get(TagName))
	//log.Tracef("all tags %v", tag)
	ApplyJSONTag(f, tag.Get(jsonTagName))
	// our own tag overrides json tag
}

/*
ApplyJSONTag uses json tag to set Name, following are list of json tag values we supported base on https://godoc.org/encoding/json
	Field int `json:"myName"`
	Field int `json:"myName,omitempty"`
However, we don't support the following:
	Field int `json:"-"`
	Field int `json:"-,"`
	Int64String int64 `json:",string"`
*/
func ApplyJSONTag(f *common.Field, tag string) {
	//log.Tracef("json tag %s", tag)
	values := strings.Split(tag, ",")
	//log.Tracef("len is %d", len(values))
	if len(values) == 0 {
		return
	}
	if len(values) > 0 && values[0] != "-" {
		f.Name = values[0]
	}
}
