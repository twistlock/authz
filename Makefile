.PHONY: all binary test image vet lint clean

SRCS = $(shell git ls-files '*.go' | grep -v '^Godeps/')
PKGS = ./core/. ./broker/. ./authz/.

default: binary

all: image
	docker build .

fmt:
	gofmt -w $(SRCS)

vet:
	@-go get -v golang.org/x/tools/cmd/vet
	$(foreach pkg,$(PKGS),go vet $(pkg);)

lint:
	@ go get -v github.com/golang/lint/golint
	$(foreach file,$(SRCS),golint $(file) || exit;)

image: test
	docker build -t twistlock/authz-broker .

binary: lint fmt vet
	CGO_ENABLED=0 go build  -o authz-broker -a -installsuffix cgo ./broker/main.go

test:  binary
	go test -v ./...

clean:
	rm authz_broker

