/*
Package fixture provides struct definition and fake data generation
*/
package fixture

import (
	"github.com/at15/go-solr/pkg/common"
	"time"
)

type JobError struct {
	Code int
	Msg  string
}

// TODO: can json & solr handle time.Duration
type Job struct {
	Name         string    `json:"name" solr:"name"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	Day          bool      `json:"day"`
	JsonIgnoreMe string    `json:"-"`
	IgnoreMe     JobError  `json:"ignore_me" solr:"-"`
	hidden       string
}

type ByteSlice struct {
	Foo []byte `solr:"bar"`
}

type JsonTag struct {
	Foo string `json:"foo"`
}

type SolrTag struct {
	Foo      string `json:"foo" solr:",type=string,docValues=true,indexed=false,stored=true,multiValued=false,required=true"`
	IgnoreMe string `solr:"-"`
}

type AllPrivate struct {
	h1 string
}

var JobFieldsSchema = []common.Field{common.Field{Name: "request_uuid", Type: "text_general", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "namespace", Type: "text_general", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "pool", Type: "text_general", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "submittedby", Type: "text_general", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "created", Type: "date", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "updated", Type: "date", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "completed", Type: "date", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "status", Type: "text_general", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "request_type", Type: "text_general", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "request_payload", Type: "text_general", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "version", Type: "text_general", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "job_error", Type: "ignored", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "dz_message", Type: "ignored", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "dz_status", Type: "ignored", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "job_link", Type: "ignored", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}, common.Field{Name: "tenant", Type: "text_general", DocValues: (*bool)(nil), Indexed: (*bool)(nil), Stored: (*bool)(nil), MultiValued: (*bool)(nil), Required: (*bool)(nil)}}
