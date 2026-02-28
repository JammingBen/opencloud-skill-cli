SERVER_URL ?= https://host.docker.internal:9200

# Get the latest tag, or "dev" if none exists
TAG := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "0.0.1")
# Get the short commit hash
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
# Check if current HEAD is exactly a tag
EXACT_TAG := $(shell git describe --tags --exact-match 2>/dev/null)

ifeq ($(EXACT_TAG),)
	# Not at a tag, append commit hash
	VERSION := $(TAG)-$(COMMIT)
else
	# At a tag, use it as is
	VERSION := $(EXACT_TAG)
endif

LDFLAGS := -ldflags "-X main.Version=$(VERSION)"

.PHONY: build
build:
	go build $(LDFLAGS) -o bin/oc-cli cmd/opencloud-cli/*.go

.PHONY: run
run:
	go run $(LDFLAGS) cmd/opencloud-cli/*.go $(ARGS)

.PHONY: login
login:
	go run $(LDFLAGS) cmd/opencloud-cli/*.go login --server-url $(SERVER_URL) --insecure

.PHONY: logout
logout:
	go run $(LDFLAGS) cmd/opencloud-cli/*.go logout

.PHONY: tidy
tidy:
	go mod tidy
