ARCH ?= $(shell uname -m)
BIN := lazysh
DESTDIR :=
GO ?= go
PKGNAME := lazysh
PREFIX := /usr/local

VERSION := $(shell git rev-list --count master)

RELEASE_DIR := ${PKGNAME}_${VERSION}_${ARCH}
PACKAGE := $(RELEASE_DIR).tar.gz
SOURCES ?= $(shell find . -name "*.go" -type f)

.PHONY: default
default: build

.PHONY: all
all: | clean release

.PHONY: clean
clean:
	$(GO) clean -i ./...
	rm -rf $(BIN) $(PKGNAME)_*

.PHONY: build
build: $(BIN)

.PHONY: release
release: $(PACKAGE)

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: install
install: build
	install -Dm755 ${BIN} $(DESTDIR)$(PREFIX)/bin/${BIN}

.PHONY: uninstall
uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/${BIN}

$(BIN): $(SOURCES)
	$(GO) build $(GOFLAGS) -ldflags '-s -w $(LDFLAGS)' $(EXTRA_GOFLAGS) -o $@

$(RELEASE_DIR):
	mkdir $(RELEASE_DIR)

$(PACKAGE): $(BIN) $(RELEASE_DIR)
	strip ${BIN}
	cp -t $(RELEASE_DIR) ${BIN}
	tar -czvf $(PACKAGE) $(RELEASE_DIR)

