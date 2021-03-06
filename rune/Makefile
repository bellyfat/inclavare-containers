CONTAINER_ENGINE := docker
GO := go

PREFIX := $(DESTDIR)/usr/local
BINDIR := $(PREFIX)/sbin
MANDIR := $(PREFIX)/share/man

GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_BRANCH_CLEAN := $(shell echo $(GIT_BRANCH) | sed -e "s/[^[:alnum:]]/-/g")
RUNC_IMAGE := runc_dev$(if $(GIT_BRANCH_CLEAN),:$(GIT_BRANCH_CLEAN))
PROJECT := github.com/opencontainers/runc
BUILDTAGS ?= seccomp selinux apparmor
COMMIT_NO := $(shell git rev-parse HEAD 2> /dev/null || true)
COMMIT ?= $(if $(shell git status --porcelain --untracked-files=no),"$(COMMIT_NO)-dirty","$(COMMIT_NO)")
VERSION := $(shell cat ./VERSION)

# TODO: rm -mod=vendor once go 1.13 is unsupported
ifneq ($(GO111MODULE),off)
	MOD_VENDOR := "-mod=vendor"
endif
ifeq ($(DEBUG),1)
	GCFLAGS=-gcflags "-N -l"
endif

PROTO_DIR := libenclave/proto libenclave/intelsgx/proto libenclave/attestation/proto
PROTOS := $(foreach dir,$(PROTO_DIR),$(patsubst %.proto,%.pb.go,$(wildcard $(dir)/*.proto)))

GO_BUILD := $(GO) build $(MOD_VENDOR) -buildmode=pie $(GCFLAGS) $(EXTRA_FLAGS) -tags "$(BUILDTAGS)" \
	-ldflags "-X main.gitCommit=$(COMMIT) -X main.version=$(VERSION) $(EXTRA_LDFLAGS)"
GO_BUILD_STATIC := CGO_ENABLED=1 $(GO) build $(MOD_VENDOR) $(GCFLAGS) $(EXTRA_FLAGS) -tags "$(BUILDTAGS) netgo osusergo" \
	-ldflags "-w -extldflags -static -X main.gitCommit=$(COMMIT) -X main.version=$(VERSION) $(EXTRA_LDFLAGS)"

.DEFAULT: rune

rune: $(PROTOS)
	$(GO_BUILD) -o rune .

%.pb.go: %.proto
	@protoc --proto_path=$(shell dirname $^) --go_out=$(shell dirname $^) $^

all: rune recvtty skeleton

recvtty:
	$(GO_BUILD) -o contrib/cmd/recvtty/recvtty ./contrib/cmd/recvtty

skeleton: libenclave/internal/runtime/pal/skeleton/liberpal-skeleton*.so
libenclave/internal/runtime/pal/skeleton/liberpal-skeleton*.so:
	make -C libenclave/internal/runtime/pal/skeleton

static: $(PROTOS)
	$(GO_BUILD_STATIC) -o rune .
	$(GO_BUILD_STATIC) -o contrib/cmd/recvtty/recvtty ./contrib/cmd/recvtty

release:
	script/release.sh -r release/$(VERSION) -v $(VERSION)

dbuild: runcimage
	$(CONTAINER_ENGINE) run $(CONTAINER_ENGINE_RUN_FLAGS) \
		--privileged --rm \
		-v $(CURDIR):/go/src/$(PROJECT) \
		$(RUNC_IMAGE) make clean all

lint:
	$(GO) vet ./...
	$(GO) fmt ./...

man:
	man/md2man-all.sh

runcimage:
	$(CONTAINER_ENGINE) build $(CONTAINER_ENGINE_BUILD_FLAGS) -t $(RUNC_IMAGE) .

test: unittest integration rootlessintegration

localtest: localunittest localintegration localrootlessintegration

unittest: runcimage
	$(CONTAINER_ENGINE) run $(CONTAINER_ENGINE_RUN_FLAGS) \
		-t --privileged --rm \
		-v /lib/modules:/lib/modules:ro \
		-v $(CURDIR):/go/src/$(PROJECT) \
		$(RUNC_IMAGE) make localunittest TESTFLAGS=$(TESTFLAGS)

localunittest: all
	$(GO) test $(MOD_VENDOR) -timeout 3m -tags "$(BUILDTAGS)" $(TESTFLAGS) -v ./...

integration: runcimage
	$(CONTAINER_ENGINE) run $(CONTAINER_ENGINE_RUN_FLAGS) \
		-t --privileged --rm \
		-v /lib/modules:/lib/modules:ro \
		-v $(CURDIR):/go/src/$(PROJECT) \
		$(RUNC_IMAGE) make localintegration TESTPATH=$(TESTPATH)

localintegration: all
	bats -t tests/integration$(TESTPATH)

rootlessintegration: runcimage
	$(CONTAINER_ENGINE) run $(CONTAINER_ENGINE_RUN_FLAGS) \
		-t --privileged --rm \
		-v $(CURDIR):/go/src/$(PROJECT) \
		-e ROOTLESS_TESTPATH \
		$(RUNC_IMAGE) make localrootlessintegration

localrootlessintegration: all
	tests/rootless.sh

shell: runcimage
	$(CONTAINER_ENGINE) run $(CONTAINER_ENGINE_RUN_FLAGS) \
		-ti --privileged --rm \
		-v $(CURDIR):/go/src/$(PROJECT) \
		$(RUNC_IMAGE) bash

install:
	install -D -m0755 rune $(BINDIR)/rune

install-bash:
	install -D -m0644 contrib/completions/bash/rune $(PREFIX)/share/bash-completion/completions/rune

install-man: man
	install -d -m 755 $(MANDIR)/man8
	install -D -m 644 man/man8/*.8 $(MANDIR)/man8

clean:
	rm -f rune rune-*
	rm -f contrib/cmd/recvtty/recvtty
	rm -rf release
	rm -rf man/man8
	rm -f $(PROTOS)
	make -C libenclave/internal/runtime/pal/skeleton clean

uninstall:
	rm -f $(BINDIR)/rune

validate:
	script/validate-gofmt
	script/validate-c
	$(GO) vet ./...

ci: validate test release

vendor:
	export GO111MODULE=on \
	$(GO) mod tidy && \
	$(GO) mod vendor && \
	$(GO) mod verify

verify-dependencies: vendor
	@test -z "$$(git status --porcelain -- go.mod go.sum vendor/)" \
		|| (echo -e "git status:\n $$(git status -- go.mod go.sum vendor/)\nerror: vendor/, go.mod and/or go.sum not up to date. Run \"make vendor\" to update"; exit 1) \
		&& echo "all vendor files are up to date."

cross: runcimage
	$(CONTAINER_ENGINE) run $(CONTAINER_ENGINE_RUN_FLAGS) \
		-e BUILDTAGS="$(BUILDTAGS)" --rm \
		-v $(CURDIR):/go/src/$(PROJECT) \
		$(RUNC_IMAGE) make localcross

localcross:
	CGO_ENABLED=1 GOARCH=arm GOARM=6 CC=arm-linux-gnueabi-gcc   $(GO_BUILD) -o runc-armel .
	CGO_ENABLED=1 GOARCH=arm GOARM=7 CC=arm-linux-gnueabihf-gcc $(GO_BUILD) -o runc-armhf .
	CGO_ENABLED=1 GOARCH=arm64 CC=aarch64-linux-gnu-gcc         $(GO_BUILD) -o runc-arm64 .
	CGO_ENABLED=1 GOARCH=ppc64le CC=powerpc64le-linux-gnu-gcc   $(GO_BUILD) -o runc-ppc64le .

package:
	$(MAKE) -C dist package

.PHONY: rune all recvtty static release dbuild lint man runcimage \
	test localtest unittest localunittest integration localintegration \
	rootlessintegration localrootlessintegration shell install install-bash \
	install-man clean uninstall validate ci \
	vendor verify-dependencies cross localcross skeleton package
