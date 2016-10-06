GOC := /usr/bin/go build
FETCHLIBS=/usr/bin/go get


BUILDDIR=$(CURDIR)/build
SRCDIR=src/noop
GOBINDIR=$(BUILDDIR)/bin
GOPATHDIR=$(BUILDDIR)/golibs

INSTALL=install
INSTALL_BIN=$(INSTALL) -m755
INSTALL_CONF=$(INSTALL) -m644

PREFIX?=$(DESTDIR)/usr
BINDIR?=$(PREFIX)/bin
SBINDIR?=$(DESTDIR)/sbin

all: noop

noop: Makefile src/noop/main.go
	mkdir -p $(GOPATHDIR) && \
	mkdir -p $(GOBINDIR) && \
	export GOPATH=$(GOPATHDIR) && \
	export GOBIN=$(GOBINDIR) && \
	cd $(SRCDIR) && \
	$(FETCHLIBS) && \
	$(GOC)

clean:
	rm -fr build/
	rm -f src/noop/noop

install:
	mkdir -p $(SBINDIR)
	$(INSTALL_BIN) src/noop $(SBINDIR)/
