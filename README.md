# go-solr

[![GoDoc](https://godoc.org/github.com/at15/go-solr?status.svg)](https://godoc.org/github.com/at15/go-solr)
[![Build Status](https://travis-ci.org/at15/go-solr.svg?branch=master)](https://travis-ci.org/at15/go-solr)
[![codebeat badge](https://codebeat.co/badges/9c885c87-c100-49ec-8414-d369cd6461f5)](https://codebeat.co/projects/github-com-at15-go-solr-master)

Solr client in golang

## Features

- Manage core via REST API
- Manage schema via REST API

Not implemented

- [ ] v2 API, use with  https://cwiki.apache.org/confluence/display/solr/v2+API
- [ ] Convert to struct (? how to say that ....)
- [ ] SolrCloud

## Roadmap

- [x] using managed schema
  - ~~generate schema xml based on golang struct (w/o?) annotation~~
  - ~~support schema less~~
- [ ] query using JSON API
- [ ] collect metrics about upstream
- [ ] client side load balancing

## Alternatives

- https://github.com/sendgrid/go-solr actively maintained
- https://github.com/vanng822/go-solr last updated 2017
- https://github.com/rtt/Go-Solr last updated 2015