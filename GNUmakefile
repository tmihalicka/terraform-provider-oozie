GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=oozie

default: build

build: fmtcheck
	go install

test: fmtcheck
	go test $(TEST) -timeout=30s -parallel=4

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)


# Currently required by tf-deploy compile
fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"


lint:
	@echo "==> Checking source code against linters..."
	@gometalinter ./$(PKG_NAME)


vendor-status:
	@govendor status


tools:
	go get -u github.com/kardianos/govendor
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: build test fmt fmtcheck lint tools vendor-status