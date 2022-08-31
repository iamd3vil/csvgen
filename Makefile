PHONY : build run fresh test clean pack-releases

BIN := csvgen.bin

HASH := $(shell git rev-parse --short HEAD)
COMMIT_DATE := $(shell git show -s --format=%ci ${HASH})
BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
VERSION := ${HASH} (${COMMIT_DATE})

STATIC := ./templates:/templates

build:
	go build -o ${BIN} -ldflags="-X 'main.buildVersion=${VERSION}' -X 'main.buildDate=${BUILD_DATE}'" ./cmd/generator/

run:
	./${BIN}

fresh: clean build

test: build
	rm -rf ./test/csvgen_gen.go
	cd test && ../csvgen.bin -file models.go -dest csvgen.go && cd ../
	go test -v ./test

clean:
	go clean
	rm -f ${BIN}
