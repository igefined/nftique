PACKAGE = $(shell go list -m)
BUILD_DIR = build
BINARY_NAME = nftique
PWD = $(shell pwd)
VERSION ?= $(shell git describe --exact-match --tags 2> /dev/null || head -1 CHANGELOG.md | cut -d ' ' -f 2)
COMMIT ?= $(shell git rev-parse HEAD)
BUILD_DATE = $(shell date +"%Y-%m-%dT%H:%M:%S")
LDFLAGS = -ldflags "-w -X ${PACKAGE}/internal/app.Version=${VERSION} -X ${PACKAGE}/internal/app.BuildDate=${BUILD_DATE} -X ${PACKAGE}/internal/app.Commit=${COMMIT}"

ifneq (,$(wildcard .env))
	include .env
	export
endif

.PHONY: update
update:
	go mod tidy
	go mod verify

bin/:
	mkdir -p bin

bin/golang-lint: | bin/
	@curl -sSfL https://raw.gcoverprofileithubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.55.2

.PHONY: lint
lint: | bin/golang-lint
	@PATH="$(realpath bin):$$PATH" golangci-lint run

.PHONY: test
test:
	@go test ./... -race -v -count=1 -coverprofile .cover
	@go tool cover -func .cover | grep "total:"

.PHONY: coverage
coverage:
	@go tool cover -html=.cover -o coverage.html

.PHONY: build
build:
	go build -tags '${TAGS}' ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME} ${PACKAGE}/cmd/${BINARY_NAME}

.PHONY: run
run: build
	./bin/nftique

.PHONY: deploy-local
deploy-local:
	cd deployment/local && docker-compose up -d