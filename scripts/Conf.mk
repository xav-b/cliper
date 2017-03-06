# Conf.mk
# vim:ft=make

# alternative : git describe --always --tags
# choose to set from here or from the command line
VERSION          ?= "0.1.1"

BUILD_PATH       := ./_dist
RELEASE_OPTS     ?= ""

SOURCES          := $(shell find . -path './vendor' -prune -o -type f -name '*.go' -print)
PACKAGES         := $(shell go list ./... | grep -v /vendor/)

GO_PROJECTS      := "/go/src/github.com/$(GIT_USER)"
GO_VERSION       := $(shell go version)
# ldflags does't support spaces in variables
CLEAN_GO_VERSION := $(shell echo "${GO_VERSION}" | sed -e 's/[^a-zA-Z0-9]/_/g')

BINARY           := ${PROJECT}
LDFLAGS          := "-X github.com/$(GIT_USER)/$(PROJECT).BuildTime=${BUILD_TIME} -X github.com/$(GIT_USER)/$(PROJECT).GoVersion=${CLEAN_GO_VERSION} -X github.com/$(GIT_USER)/$(PROJECT).GitCommit=${GIT_COMMIT}"

