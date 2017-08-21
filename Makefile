PKG = github.com/at15/go-solr
VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
BUILD_USER = ${USER}
FLAGS = -X main.Version=$(VERSION) -X main.BuildCommit=$(BUILD_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildUser=$(BUILD_USER)

.PHONY: example install test test-in-docker docker-test

example:
	cd example/job; go run main.go
install:
	rm -f $(shell which solrgo)
	go install  -ldflags "$(FLAGS)" ./cmd/solrgo
test:
	go test -v -cover ./pkg/...
test-in-docker:
	./script/wait-for-it.sh solr:8983
	make install
	solrgo core create demo
	go test -v -cover ./pkg/...
	cd example/job; go run main.go
docker-test:
	cd script; docker-compose run golang
	cd script; docker-compose down
fmt:
	gofmt -d -l -w ./pkg
	gofmt -d -l -w ./cmd/solrgo
	gofmt -d -l -w ./example/job
loc:
	cloc --exclude-dir=vendor,.idea,script .