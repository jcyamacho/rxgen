TOOLS := $(CURDIR)/.tools

tools:
	mkdir -p ${TOOLS}
	GOBIN=${TOOLS} go install github.com/caarlos0/svu@latest
	GOBIN=${TOOLS} go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	GOBIN=${TOOLS} go install github.com/goreleaser/goreleaser@latest

lint: tools
	${TOOLS}/golangci-lint run

build: tools
	${TOOLS}/goreleaser release --snapshot --rm-dist

ARG ?= next
VERSION := $(shell ${TOOLS}/svu $(ARG))
release: tools
	git tag "$(VERSION)"
	git push origin "$(VERSION)"

release-major: ARG=major
release-major: release

release-minor: ARG=minor
release-minor: release

release-patch: ARG=patch
release-patch: release
