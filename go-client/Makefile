NAME=go-client
BUILDDIR=build

BASEPATH := $(shell pwd)
CLIENTSRCFILE := $(BASEPATH)/cmd/awaken/

LDFLAGS=-w -s

CLIENTBUILD=CGO_ENABLED=0 go build -trimpath

all:
	GOARCH=amd64 GOOS=darwin $(CLIENTBUILD) -ldflags "$(LDFLAGS)" -o $(BUILDDIR)/darwin/KiwiClient $(CLIENTSRCFILE)

	GOARCH=amd64 GOOS=windows $(CLIENTBUILD) -ldflags "$(LDFLAGS) -H windowsgui" -o $(BUILDDIR)/windows/KiwiClient.exe $(CLIENTSRCFILE)

	GOARCH=amd64 GOOS=linux $(CLIENTBUILD) -ldflags "$(LDFLAGS)" -o $(BUILDDIR)/linux-amd64/KiwiClient $(CLIENTSRCFILE)

	GOARCH=arm64 GOOS=linux $(CLIENTBUILD) -ldflags "$(LDFLAGS)" -o $(BUILDDIR)/linux-arm64/KiwiClient $(CLIENTSRCFILE)

dev:
	go build -o $(BUILDDIR)/windows/KiwiClient.exe $(CLIENTSRCFILE)