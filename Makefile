.PHONY: example install test test-in-docker docker-test

example:
	cd example/job; go run main.go
install:
	go install ./cmd/gosolr
test:
	go test -v -cover ./pkg/...
test-in-docker:
	./script/wait-for-it.sh solr:8983
	go test -v -cover ./pkg/...
docker-test:
	cd script; docker-compose run golang
	cd script; docker-compose down
fmt:
	gofmt -d -l -w ./pkg
	gofmt -d -l -w ./cmd/gosolr
loc:
	cloc --exclude-dir=vendor,.idea,script .