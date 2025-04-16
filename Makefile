# Makefile

BINARY_NAME = tkeyauth
MAIN = ./cmd/main.go

OS := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

.PHONY: all build clean

all: build

build:
	@echo "Building for $(OS)/$(ARCH)..."
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -o $(BINARY_NAME) $(MAIN)
	@echo "Binary created at ./$(BINARY_NAME)"

clean:
	@echo "Cleaning build output..."
	@rm -f $(BINARY_NAME)
