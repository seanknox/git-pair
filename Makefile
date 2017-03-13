GO        ?= go
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   :=
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin
APP       := git-pair
OUT	 			:= ./bin/$(APP)

.PHONY: build
build:
	$(GO) build -ldflags "$(LDFLAGS)" -o $(OUT) cmd/git_pair.go

.PHONY: clean
clean:
	rm -frv $(OUT)
	go clean

.PHONY: test
test:
	ginkgo --skipPackage vendor/ -r

HAS_GINKGO := $(shell command -v ginkgo;)
HAS_GIT := $(shell command -v git;)

.PHONY: bootstrap
bootstrap:
ifndef HAS_GINKGO
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega
endif
ifndef HAS_GIT
	$(error You must install Git)
endif
