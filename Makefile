GOPATH ?= $(HOME)/go
GOBIN ?= $(GOPATH)/bin
GOIMPORTS = $(GOBIN)/goimports

PKG = ./addon/... \
      ./cmd/... \
      ./command/... \
      ./nas/... \
      ./scm/... \
      ./ssh/...

PKGDIR = $(subst /...,,$(PKG))

all: cmd

fmt: $(GOIMPORTS)
	$(GOIMPORTS) -w $(PKGDIR)

vet:
	go vet $(PKG)

cmd: fmt vet
	go build -ldflags="-w -s" -o bin/addon github.com/konveyor/tackle2-addon/cmd

# Run unit tests.
test: cmd
	go test -count=1 -v $(PKG)

# Ensure goimports installed.
$(GOIMPORTS):
	go install golang.org/x/tools/cmd/goimports@v0.24
