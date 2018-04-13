.PHONY: format
format:
	@find . -type f -name "*.go*" -print0 | xargs -0 gofmt -s -w

.PHONY: debs
debs:
	GOPATH=$(GOPATH) go get ./...
	GOPATH=$(GOPATH) go get -u gopkg.in/check.v1
	GOPATH=$(GOPATH) go get -u github.com/golang/lint/golint
	GOPATH=$(GOPATH) go get -u github.com/fortytw2/leaktest

.PHONY: test
test:
	GOPATH=$(GOPATH) golint -set_exit_status $(go list ./...)
	GOPATH=$(GOPATH) go test -race -gcflags -m

.PHONY: bench
bench:
	GOPATH=$(GOPATH) go test -bench=. -check.b -benchmem

# Clean junk
.PHONY: clean
clean:
	GOPATH=$(GOPATH) go clean ./...
