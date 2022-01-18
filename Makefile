VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=blog \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=blogd \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

build-local-linux: go.sum generate
	env GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./build/checkersd $(BUILD_FLAGS) ./cmd/checkersd

generate:
ifeq ($(GENERATE),1)
	go generate ./...
endif

