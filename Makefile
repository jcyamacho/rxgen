TOOLS := $(CURDIR)/.tools

tools:
	@mkdir -p ${TOOLS}
	@GOBIN=${TOOLS} go install github.com/caarlos0/svu@latest
	@GOBIN=${TOOLS} go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@GOBIN=${TOOLS} go install github.com/goreleaser/goreleaser@latest

lint: tools
	${TOOLS}/golangci-lint run

build: tools
	${TOOLS}/goreleaser release --snapshot --rm-dist

VERSION ?= $(shell ${TOOLS}/svu next)
release: tools
	@echo "Release $(VERSION)"
	git tag "$(VERSION)"
	git push origin "$(VERSION)"

release-major: VERSION=$(shell ${TOOLS}/svu major)
release-major: release

release-minor: VERSION=$(shell ${TOOLS}/svu minor)
release-minor: release

release-patch: VERSION=$(shell ${TOOLS}/svu patch)
release-patch: release
