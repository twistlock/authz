.PHONY: all binary image clean

default: binary

all: image
	docker build .

image: binary
	docker build .

binary:
	CGO_ENABLED=0 go build  -o twistlock_authz_plugin -a -installsuffix cgo ./broker/main.go

clean:
	rm twistlock_authz_plugin