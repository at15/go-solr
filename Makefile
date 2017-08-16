.PHONY: test

test:
	go test -v -cover ./pkg/...
fmt:
	gofmt -d -l -w ./pkg
loc:
	cloc --exclude-dir=vendor,.idea,script .