
GO_VERSION=1.17
ALL_GO_FILES=$(shell find . -type f  -name '*.go')

tidy:
	go mod tidy -compat=$(GO_VERSION)

clean:
	go clean ./...

fmt:
	goimports -w $(ALL_GO_FILES)

lint:
	golint ./...
