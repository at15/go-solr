# go-solr

[![GoDoc](https://godoc.org/github.com/at15/go-solr?status.svg)](https://godoc.org/github.com/at15/go-solr)
[![Build Status](https://travis-ci.org/at15/go-solr.svg?branch=master)](https://travis-ci.org/at15/go-solr)
[![codebeat badge](https://codebeat.co/badges/9c885c87-c100-49ec-8414-d369cd6461f5)](https://codebeat.co/projects/github-com-at15-go-solr-master)

Solr client in golang

## Usage

- install the cli using `go get -u github.com/at15/go-solr/cmd/solrgo`, it can create core and index json document

````go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/at15/go-solr/solr"
)

const coreName = "job"

func main() {
    c := solr.Config{}
    solrClient, err := solr.NewClient(c)
    if err != nil {
        log.Fatal(err)
        return
    }
    if err := solrClient.IsUp(context.Background()); err != nil {
        log.Fatalf("Solr is not up %v", err)
        return
    }
    log.Println("Solr is up")
    solrClient.UseCore(coreName)
    if status, err := solrClient.DefaultCore.Status(context.Background(), false); err != nil {
        log.Fatalf("Check core status failed %v", err)
        return
    } else {
        log.Printf("Got status for core %s %v\n", coreName, status)
    }
}
````

## Features

- standalone command line util (single binary)
- query builder
- auto convert struct to JSON when ingest document 

Not implemented

- [ ] v2 API https://lucene.apache.org/solr/guide/6_6/v2-api.html
- [ ] SolrCloud (zookeeper aware)

## Roadmap

- [x] using managed schema
  - ~~generate schema xml based on golang struct (w/o?) annotation~~
  - ~~support schema less~~
- [x] query using JSON API
- [ ] v2 API
- [ ] collect metrics about upstream
- [ ] client side load balancing

## Alternatives

- https://github.com/sendgrid/go-solr actively maintained, support zk
- https://github.com/vanng822/go-solr last updated 2017
- https://github.com/rtt/Go-Solr last updated 2015