# Go.mk
# vim:ft=make

PLATFORM ?= "windows linux darwin"

all: $(BINARY)

container-up: ## start a pending container workspace
	docker run -d --name $(PROJECT) \
		-v $(PWD):$(GO_PROJECTS)/$(PROJECT) \
		-w $(GO_PROJECTS)/$(PROJECT) $(CONTAINER) sleep infinity

container-shell: ## start a shell in the workspace container
	docker exec -it $(PROJECT) bash

crossbuild: $(SOURCES) ## cross-compile binary to windows linux and osx
	CC="gcc" gox -verbose -cgo \
		-ldflags $(LDFLAGS) \
		-os=$(PLATFORM) \
		-arch="amd64" \
		-output="$(BUILD_PATH)/$(VERSION)/{{.Dir}}-{{.OS}}-{{.Arch}}" .

# usage:
# PLATFORM=darwin \
# COMMENT="Buggy MVP completed" \
# RELEASE_OPTS="-prerelease -b 'Buggy MVP completed'" \
# make release
release: crossbuild ## git tag and publish a release on Github
ifndef COMMENT
	$(error no tag description provided)
endif
	git tag -a $(VERSION) -m '$(COMMENT)'
	git push --tags
	ghr $(RELEASE_OPTS) v$(VERSION) $(BUILD_PATH)/$(VERSION)/

$(BINARY): $(SOURCES) ## compile project
	go build -v -ldflags ${LDFLAGS} -o ${BINARY}

.PHONY: install.tools
install-tools: ## install development tools
	# code coverage
	go get github.com/axw/gocov/gocov
	# cross-compilation
	go get github.com/mitchellh/gox
	# github release publication
	go get github.com/tcnksm/ghr
	# code linting
	# FIXME make circleci to fail
	go get github.com/alecthomas/gometalinter && \
		gometalinter --install --update

install-hack: install.tools ## install dev tools and project deps
	go get ./...

install: ## compile and globally install the cli
	go install -ldflags ${LDFLAGS}

lint: ## lint code with various tools
	test -z "$(go fmt ./...)"
	GO_VENDOR=1 gometalinter --deadline=25s ./...

test: ## go test code
	go test ./... $(TESTARGS)

.PHONY: godoc
godoc: ## run Go doc server
	godoc -http=0.0.0.0:6060

.PHONY: clean
clean: ## remove build artifacts
	[[ -d ${BUILD_PATH} ]] && rm -rf ${BUILD_PATH}
	[[ -f ${BINARY} ]] && rm -rf ${BINARY}
