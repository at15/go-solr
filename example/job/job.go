package main

import "time"

/*

cd bin
./solr start -f -v
# this created managed config
./solr create -c jobs


 */
type Job struct {
	Id        string
	Namespace string
	Pool      string
	Created   time.Time
	Updated   time.Time
	Type      int
}
