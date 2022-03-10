# PREFIX is environment variable, but if it is not set, then set default value
ifeq ($(PREFIX),)
    PREFIX := /usr/local/bin
endif

compile:
	go build gotextme.go

copy: gotextme
	sudo install -d $(PREFIX)
	sudo install -m 755 mquery $(PREFIX)

install: compile copy
