/*
Package fixture provides struct definitions for testing InferSchema in schema package
*/
package fixture

import "time"

type Job struct {
	Name      string    `json:"name" solr:"foo"`
	StartTime time.Time `json:"startTime"`
	Day       bool      `json:"day"`
	IgnoreMe  string    `json:"ignore_me" solr:"-"`
	hidden    string
}

type ByteSlice struct {
	Foo []byte `solr:"bar"`
}

type AllPrivate struct {
	h1 string
	h2 string
	h3 string
}
