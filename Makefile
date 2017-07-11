SHELL := /bin/bash
.DEFAULT_GOAL := all
MODULE_NAME=$(shell basename `pwd`)
ARCH := amd64
GPDB_VERSION := $(shell ../../../../getversion --short)

export GOPATH := $(shell dirname `pwd`)/..
export PATH := $(PATH):$(GOPATH)/bin

all : build test

dependencies :
		go get github.com/greenplum-db/gpbackup/utils
		go get github.com/greenplum-db/gpbackup/testutils
		go get github.com/cppforlife/go-semi-semantic/version
		go get github.com/onsi/ginkgo/ginkgo
		go get golang.org/x/tools/cmd/goimports
		go get github.com/onsi/gomega
		go get github.com/jessevdk/go-flags
# Counterfeiter is not a proper dependency of the app. It is only used occasionally to generate a test class that
# is then checked in.  At the time of that generation, it can be added back to run the dependency list, temporarily.
#		go get github.com/maxbrunsfeld/counterfeiter

format : dependencies
		goimports -w .
		go fmt .

unit : dependencies sshd_build
		ginkgo -r -randomizeSuites -randomizeAllSpecs -race --skipPackage=integrations

sshd_build : dependencies
		make -C integrations/sshd

integration: dependencies sshd_build unit
		ginkgo -r -randomizeAllSpecs -race integrations

test : format unit sshd_build integration

push : format
		git pull -r && make test && git push

build : dependencies
		go build -ldflags "-X gp_upgrade/commands.GpdbVersion=$(GPDB_VERSION)" -o $(GOPATH)/bin/$(MODULE_NAME)

coverage: dependencies format sshd_build build 
		./scripts/run_coverage.sh


linux :
		GOOS=$@ GOARCH=$(ARCH) go build -ldflags "-X gp_upgrade/commands.GpdbVersion=$(GPDB_VERSION)" -o $(GOPATH)/bin/$(MODULE_NAME).$@
darwin :
		GOOS=$@ GOARCH=$(ARCH) go build -ldflags "-X gp_upgrade/commands.GpdbVersion=$(GPDB_VERSION)" -o $(GOPATH)/bin/$(MODULE_NAME).$@

platforms: linux darwin
