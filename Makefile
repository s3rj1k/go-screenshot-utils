GO_BIN ?= go
CURL_BIN ?= curl
SHELL_BIN ?= sh

deps: check-gopath
	$(GO_BIN) get -u github.com/s3rj1k/go-screenshots
	$(GO_BIN) get -u golang.org/x/net/publicsuffix
	$(GO_BIN) get -u github.com/gin-gonic/gin
	$(GO_BIN) get -u github.com/gin-contrib/cors

	# linters
	$(CURL_BIN) -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | $(SHELL_BIN) -s -- -b ${GOPATH}/bin v1.16.0
	$(GO_BIN) get -u github.com/Quasilyte/go-consistent

lint:
	golangci-lint run
	go-consistent -v ./...

check-gopath:
ifndef GOPATH
	$(error GOPATH is undefined)
endif
