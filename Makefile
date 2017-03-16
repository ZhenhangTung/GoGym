GOGYM_PKGS = $$(go list ./... | grep -v /vendor/)

check-path:
ifndef GOPATH
	@echo "FATAL: you must declare GOPATH environment variable, for more"
	@echo "       details, please check"
	@echo "       http://golang.org/doc/code.html#GOPATH"
	@exit 1
endif
ifneq ($(subst ~,$(HOME),$(GOPATH))/src/github.com/ZhenhangTung/GoGym, $(PWD))
	@echo "FATAL: you must clone tsuru inside your GOPATH To do so,"
	@echo "       you can run go get github.com/ZhenhangTung/GoGym/..."
	@echo "       or clone it manually to the dir $(GOPATH)/src/github.com/ZhenhangTung/GoGym"
	@exit 1
endif
	@exit 0

all: check-path test race install

install:
	gp install $(GO_EXTRAFLAGS) $(GOGYM_PKGS) $$(go list ./... | grep -v /vendor/)

test:
	go clean $(GO_EXTRAFLAGS) $(GOGYM_PKGS)
	go test $(GO_EXTRAFLAGS) $(GOGYM_PKGS)

race:
	go test $(GO_EXTRAFLAGS) -race -i $(GOGYM_PKGS)
	go test $(GO_EXTRAFLAGS) -race $(GOGYM_PKGS)
