# PREFIX is environment variable, but if it is not set, then set default value
ifeq ($(PREFIX),)
    PREFIX := /usr/bin
endif

compile:
	go build gotextme.go

install: compile
	sudo install -d $(PREFIX)
	sudo install -m 755 gotextme $(PREFIX)
