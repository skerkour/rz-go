.PHONY: all test release build benchmarks

VERSION := $(shell cat version.go| grep "\sVersion =" | cut -d '"' -f2)

all: test build

test:
	go tool vet -all -shadowstrict .
	go test -v -race -bench . -benchmem ./...

build:
	go build ./...

release:
	git tag v$(VERSION)
	git push origin v$(VERSION)

benchmarks:
	cd benchmarks && ./run.sh
