ifndef GOOS
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	GOOS := darwin
else ifeq ($(UNAME_S),Linux)
	GOOS := linux
else
$(error "$$GOOS is not defined. If you are using Windows, try to re-make using 'GOOS=windows make ...' ")
endif
endif

PACKAGES    := $(shell go list ./... | grep -v '/vendor/' | grep -v '/crypto/ed25519/chainkd' | grep -v '/mining/tensority')
PACKAGES += 'core/mining/tensority/go_algorithm'

BUILD_FLAGS := -ldflags "-X core/version.GitCommit=`git rev-parse HEAD`"

#MINER_BINARY32 := miner-$(GOOS)_386
MINER_BINARY64 := miner-$(GOOS)_amd64

#EIYAROD_BINARY32 := eiyarod-$(GOOS)_386
EIYAROD_BINARY64 := eiyarod-$(GOOS)_amd64

#EIYAROCLI_BINARY32 := eiyarocli-$(GOOS)_386
EIYAROCLI_BINARY64 := eiyarocli-$(GOOS)_amd64

VERSION := $(shell awk -F= '/Version =/ {print $$2}' version/version.go | tr -d "\" ")

#MINER_RELEASE32 := miner-$(VERSION)-$(GOOS)_386
MINER_RELEASE64 := miner-$(VERSION)-$(GOOS)_amd64

#EIYAROD_RELEASE32 := eiyarod-$(VERSION)-$(GOOS)_386
EIYAROD_RELEASE64 := eiyarod-$(VERSION)-$(GOOS)_amd64

#EIYAROCLI_RELEASE32 := eiyarocli-$(VERSION)-$(GOOS)_386
EIYAROCLI_RELEASE64 := eiyarocli-$(VERSION)-$(GOOS)_amd64

#EIYARO_RELEASE32 := eiyaro-$(VERSION)-$(GOOS)_386
EIYARO_RELEASE64 := eiyaro-$(VERSION)-$(GOOS)_amd64

all: test target release-all install

eiyarod:
	@echo "Building eiyarod to cmd/eiyarod/eiyarod"
	@go build $(BUILD_FLAGS) -o cmd/eiyarod/eiyarod cmd/eiyarod/main.go

eiyarod-simd:
	@echo "Building SIMD version eiyarod to cmd/eiyarod/eiyarod"
	@cd mining/tensority/cgo_algorithm/lib/ && make
	@go build -tags="simd" $(BUILD_FLAGS) -o cmd/eiyarod/eiyarod cmd/eiyarod/main.go

eiyarocli:
	@echo "Building eiyarocli to cmd/eiyarocli/eiyarocli"
	@go build $(BUILD_FLAGS) -o cmd/eiyarocli/eiyarocli cmd/eiyarocli/main.go

install:
	@echo "Installing eiyarod and eiyarocli to $(GOPATH)/bin"
	@go install ./cmd/eiyarod
	@go install ./cmd/eiyarocli

target:
	mkdir -p $@

binary: target/$(EIYAROD_BINARY64) target/$(EIYAROCLI_BINARY64) target/$(MINER_BINARY64)

ifeq ($(GOOS),windows)
release: binary
	cd target && cp -f $(MINER_BINARY64) $(MINER_BINARY64).exe
	cd target && cp -f $(EIYAROD_BINARY64) $(EIYAROD_BINARY64).exe
	cd target && cp -f $(EIYAROCLI_BINARY64) $(EIYAROCLI_BINARY64).exe
	cd target && md5sum $(MINER_BINARY64).exe $(EIYAROD_BINARY64).exe $(EIYAROCLI_BINARY64).exe >$(EIYARO_RELEASE64).md5
	cd target && zip $(EIYARO_RELEASE64).zip $(MINER_BINARY64).exe $(EIYAROD_BINARY64).exe $(EIYAROCLI_BINARY64).exe $(EIYARO_RELEASE64).md5
	cd target && rm -f $(MINER_BINARY64) $(EIYAROD_BINARY64) $(EIYAROCLI_BINARY64) $(MINER_BINARY64).exe $(EIYAROD_BINARY64).exe $(EIYAROCLI_BINARY64).exe $(EIYARO_RELEASE64).md5
else
release: binary
	cd target && md5sum $(MINER_BINARY64) $(EIYAROD_BINARY64) $(EIYAROCLI_BINARY64) >$(EIYARO_RELEASE64).md5
	cd target && tar -czf $(EIYARO_RELEASE64).tgz $(MINER_BINARY64) $(EIYAROD_BINARY64) $(EIYAROCLI_BINARY64) $(EIYARO_RELEASE64).md5
	cd target && rm -f $(MINER_BINARY64) $(EIYAROD_BINARY64) $(EIYAROCLI_BINARY64) $(EIYARO_RELEASE64).md5
endif

release-all: clean
#	GOOS=darwin  make release
	GOOS=linux   make release
	GOOS=windows make release

clean:
	@echo "Cleaning binaries built..."
	@rm -rf cmd/eiyarod/eiyarod
	@rm -rf cmd/eiyarocli/eiyarocli
	@rm -rf cmd/miner/miner
	@rm -rf target
	@rm -rf $(GOPATH)/bin/eiyarod
	@rm -rf $(GOPATH)/bin/eiyarocli
	@echo "Cleaning temp test data..."
	@rm -rf test/pseudo_hsm*
	@rm -rf blockchain/pseudohsm/testdata/pseudo/
	@echo "Cleaning sm2 pem files..."
	@rm -rf crypto/sm2/*.pem
	@echo "Done."

target/$(EIYAROD_BINARY64):
	CGO_ENABLED=0 GOARCH=amd64 go build $(BUILD_FLAGS) -o $@ cmd/eiyarod/main.go

target/$(EIYAROCLI_BINARY64):
	CGO_ENABLED=0 GOARCH=amd64 go build $(BUILD_FLAGS) -o $@ cmd/eiyarocli/main.go

target/$(MINER_BINARY64):
	CGO_ENABLED=0 GOARCH=amd64 go build $(BUILD_FLAGS) -o $@ cmd/miner/main.go

test:
	@echo "====> Running go test"
	@go test -tags "network" $(PACKAGES)

benchmark:
	@go test -bench $(PACKAGES)

functional-tests:
	@go test -timeout=5m -tags="functional" ./test

ci: test functional-tests

.PHONY: all target release-all clean test benchmark
