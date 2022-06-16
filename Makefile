#!/usr/bin/make
TOOL := pdnsclient
SHELL := /bin/bash
export PATH = /usr/sbin:/bin:/sbin:/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/build/bin:/home/runner/go:/home/runner/go/bin:/snap/bin:~/go/bin
TEST_DIRS := ./...

lint:
ifdef INTERACTIVE
	golangci-lint run -v $(TEST_DIRS)
else
	golangci-lint run --out-format checkstyle -v $(TEST_DIRS) 1> reports/checkstyle-lint.xml
endif
.PHONY: lint

deps:
	GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2

test:
	@echo ""
	@echo "***** Testing ${TOOL} *****"
	go test -a -v -race -coverprofile=reports/coverage.txt -covermode=atomic -json ./... 1> reports/testreport.json
	@echo ""

