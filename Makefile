.PHONY: test docker-test

test:
	go test -v -cover ./pkg/...
	cd example/job; go run main.go
docker-test:
	cd script; docker-compose run golang
	cd script; docker-compose down
fmt:
	gofmt -d -l -w ./pkg
loc:
	cloc --exclude-dir=vendor,.idea,script .