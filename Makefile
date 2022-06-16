#!/usr/bin/make
SHELL := /bin/bash
export PATH = /usr/sbin:/bin:/sbin:/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/build/bin:/home/runner/go:/home/runner/go/bin:/snap/bin:~/go/bin


lint:
ifdef INTERACTIVE
	golangci-lint run -v $(TEST_DIRS)
else
	golangci-lint run --out-format checkstyle -v $(TEST_DIRS) 1> reports/checkstyle-lint.xml
endif
.PHONY: lint

deps:
	GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
