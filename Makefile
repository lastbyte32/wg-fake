APPNAME = wg-fake
OUTSUFFIX = build/$(APPNAME)
VERSION := $(shell git describe)
BUILDOPTS = -a -tags netgo
LDFLAGS = -ldflags '-s -w -extldflags "-static" -X main.version=$(VERSION)'
LDFLAGS_NATIVE = -ldflags '-s -w -X main.version=$(VERSION)'
MAIN_PACKAGE = ./cmd/$(APPNAME)


GO := go

src = $(wildcard *.go */*.go */*/*.go)

native: bin-native
all: bin-linux-amd64 bin-linux-386 bin-linux-arm bin-linux-arm64 \
	bin-linux-mips bin-linux-mipsle bin-linux-mips64 bin-linux-mips64le \
	bin-darwin-amd64 bin-darwin-arm64 \
	bin-windows-amd64 bin-windows-386 bin-windows-arm

bin-native: $(OUTSUFFIX)
bin-linux-amd64: $(OUTSUFFIX).linux-amd64
bin-linux-386: $(OUTSUFFIX).linux-386
bin-linux-arm: $(OUTSUFFIX).linux-arm
bin-linux-arm64: $(OUTSUFFIX).linux-arm64
bin-linux-mips: $(OUTSUFFIX).linux-mips
bin-linux-mipsle: $(OUTSUFFIX).linux-mipsle
bin-linux-mips64: $(OUTSUFFIX).linux-mips64
bin-linux-mips64le: $(OUTSUFFIX).linux-mips64le
bin-darwin-amd64: $(OUTSUFFIX).darwin-amd64
bin-darwin-arm64: $(OUTSUFFIX).darwin-arm64
bin-windows-amd64: $(OUTSUFFIX).windows-amd64.exe
bin-windows-386: $(OUTSUFFIX).windows-386.exe
bin-windows-arm: $(OUTSUFFIX).windows-arm.exe

$(OUTSUFFIX): $(src)
	$(GO) build $(LDFLAGS_NATIVE) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).linux-amd64: $(src)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).linux-386: $(src)
	CGO_ENABLED=0 GOOS=linux GOARCH=386 $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).linux-arm: $(src)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).linux-arm64: $(src)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).linux-mips: $(src)
	CGO_ENABLED=0 GOOS=linux GOARCH=mips GOMIPS=softfloat $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).linux-mips64: $(src)
	CGO_ENABLED=0 GOOS=linux GOARCH=mips64 GOMIPS=softfloat $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).linux-mipsle: $(src)
	CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).linux-mips64le: $(src)
	CGO_ENABLED=0 GOOS=linux GOARCH=mips64le GOMIPS=softfloat $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).darwin-amd64: $(src)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).darwin-arm64: $(src)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).windows-amd64.exe: $(src)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).windows-386.exe: $(src)
	CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)

$(OUTSUFFIX).windows-arm.exe: $(src)
	CGO_ENABLED=0 GOOS=windows GOARCH=arm GOARM=7 $(GO) build $(BUILDOPTS) $(LDFLAGS) -o $@ $(MAIN_PACKAGE)


clean:
	rm -f build/*

run:
	$(GO) run $(LDFLAGS) $(MAIN_PACKAGE)

install:
	$(GO) install $(LDFLAGS_NATIVE) $(MAIN_PACKAGE)

.PHONY: clean all native fmt install \
	bin-native \
	bin-linux-amd64 \
	bin-linux-386 \
	bin-linux-arm \
	bin-linux-arm64 \
	bin-linux-mips \
	bin-linux-mipsle \
	bin-linux-mips64 \
	bin-linux-mips64le \
	bin-darwin-amd64 \
	bin-darwin-arm64 \
	bin-windows-amd64 \
	bin-windows-386 \
	bin-windows-arm \