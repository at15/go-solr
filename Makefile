PKG = github.com/at15/go-solr
VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
BUILD_USER = ${USER}
FLAGS = -X main.Version=$(VERSION) -X main.BuildCommit=$(BUILD_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildUser=$(BUILD_USER)

.PHONY: example install test test-in-docker docker-test godoc

example:
	cd example/job; go run main.go
install:
	rm -f $(shell which solrgo)
	go install  -ldflags "$(FLAGS)" ./cmd/solrgo
test:
	go test -v -cover ./solr/...
test-in-docker:
	./script/wait-for-it.sh solr:8983
	make install
	solrgo core create demo
	solrgo core create film --configSet film
	solrgo core create job --configSet job
	cd example/film; solrgo core index film films.json
	go test -v -cover ./solr/...
docker-test:
	cd script; docker-compose build
	cd script; docker-compose run golang
	cd script; docker-compose down
fmt:
	gofmt -d -l -w ./solr
	gofmt -d -l -w ./cmd/solrgo
	gofmt -d -l -w ./example/job
loc:
	cloc --exclude-dir=vendor,.idea,script .
godoc:
	@echo open http://localhost:6060/pkg/github.com/at15/go-solr
	godoc -http=":6060"
