/*
Package fixture provides struct definition and fake data generation
*/
package fixture

import "time"

// TODO: can json & solr handle time.Duration
type Job struct {
	Name         string    `json:"name" solr:"name"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	Day          bool      `json:"day"`
	JsonIgnoreMe string    `json:"-"`
	IgnoreMe     string    `json:"ignore_me" solr:"-"`
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
