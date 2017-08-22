/*
Package fixture provides struct definitions for testing InferSchema in schema package
*/
package fixture

import "time"

// TODO: move some test like tag etc. out to ApplyTag
type Job struct {
	Name      string    `json:"name" solr:"name2"`
	StartTime time.Time `json:"startTime"`
	Day       bool      `json:"day"`
	IgnoreMe  string    `json:"ignore_me" solr:"-"`
	hidden    string
}

type ByteSlice struct {
	Foo []byte `solr:"bar"`
}

type JsonTag struct {
	Foo string `json:"foo"`
}

type SolrTag struct {
	Foo string `json:"foo" solr:",type=string,docValues=true,indexed=false,stored=true,multiValued=false,required=true"`
}

type AllPrivate struct {
	h1 string
	h2 string
	h3 string
}
