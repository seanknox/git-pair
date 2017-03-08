GO        ?= go
PKG       := $(shell glide novendor)
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   :=
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin
APP       := git-pair

.PHONY: build
build:
	$(GO) build -ldflags $(LDFLAGS) -v -o $(BINDIR)/$(APP)

.PHONY: clean
clean:
	rm -frv $(OUT)
	go clean

.PHONY: test
test:
	./bin/ginkgo --skipPackage vendor/ -r

HAS_GLIDE := $(shell command -v glide;)
HAS_GOX := $(shell command -v gox;)
HAS_GIT := $(shell command -v git;)

.PHONY: bootstrap
bootstrap:
ifndef HAS_GLIDE
	go get -u github.com/Masterminds/glide
endif
ifndef HAS_GOX
	go get -u github.com/mitchellh/gox
endif

ifndef HAS_GIT
	$(error You must install Git)
endif
	glide install --strip-vendor
	go build -o bin/ginkgo ./vendor/github.com/onsi/ginkgo/ginkgo
