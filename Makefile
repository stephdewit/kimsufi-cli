exe = kimsufi
PREFIX ?= /usr
dest = $(DESTDIR)$(PREFIX)/bin/$(exe)
INSTALL = install

.PHONY: all clobber install install-strip uninstall

all: $(exe)

$(exe): *.go availability/*.go
	go build -o $(exe) .

clobber:
	rm -vf $(exe)

$(dest): $(exe)
	$(INSTALL) -vD $(exe) $(dest)

install: $(dest)

install-strip:
	$(MAKE) INSTALL='$(INSTALL) -s' install

uninstall:
	rm -vf $(dest)
