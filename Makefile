# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: ronin android ios ronin-cross evm all test clean bootnode
.PHONY: ronin-linux ronin-linux-386 geth-linux-amd64 geth-linux-mips64 geth-linux-mips64le
.PHONY: ronin-linux-arm ronin-linux-arm-5 geth-linux-arm-6 geth-linux-arm-7 geth-linux-arm64
.PHONY: ronin-darwin ronin-darwin-386 geth-darwin-amd64
.PHONY: ronin-windows ronin-windows-386 geth-windows-amd64

GOBIN = ./build/bin
GO ?= latest
GORUN = go run
RONIN_CONTRACTS_PATH = ../ronin-dpos-contracts
RONIN_CONTRACTS_OUTPUT_PATH = ./tmp/contracts
GEN_CONTRACTS_OUTPUT_PATH = ./consensus/consortium/generated_contracts

generate-contract:
	@echo "Generating"
	solc --abi --bin $(RONIN_CONTRACTS_PATH)/contracts/ronin/staking/Staking.sol -o $(RONIN_CONTRACTS_OUTPUT_PATH)/staking --include-path $(RONIN_CONTRACTS_PATH)/node_modules/ --base-path $(RONIN_CONTRACTS_PATH) --overwrite --optimize
	abigen --abi $(RONIN_CONTRACTS_OUTPUT_PATH)/staking/Staking.abi --bin $(RONIN_CONTRACTS_OUTPUT_PATH)/staking/Staking.bin --pkg staking --out $(GEN_CONTRACTS_OUTPUT_PATH)/staking/staking.go

	solc --abi --bin $(RONIN_CONTRACTS_PATH)/contracts/ronin/validator/RoninValidatorSet.sol -o $(RONIN_CONTRACTS_OUTPUT_PATH)/validator --include-path $(RONIN_CONTRACTS_PATH)/node_modules/ --base-path $(RONIN_CONTRACTS_PATH) --overwrite --optimize
	abigen --abi $(RONIN_CONTRACTS_OUTPUT_PATH)/validator/RoninValidatorSet.abi --bin $(RONIN_CONTRACTS_OUTPUT_PATH)/validator/RoninValidatorSet.bin --pkg roninValidatorSet --out $(GEN_CONTRACTS_OUTPUT_PATH)/ronin_validator_set/ronin_validator_set.go

	solc --abi --bin $(RONIN_CONTRACTS_PATH)/contracts/ronin/SlashIndicator.sol -o $(RONIN_CONTRACTS_OUTPUT_PATH)/slashing --include-path $(RONIN_CONTRACTS_PATH)/node_modules/ --base-path $(RONIN_CONTRACTS_PATH) --overwrite --optimize
	abigen --abi $(RONIN_CONTRACTS_OUTPUT_PATH)/slashing/SlashIndicator.abi --bin $(RONIN_CONTRACTS_OUTPUT_PATH)/slashing/SlashIndicator.bin --pkg slashIndicator --out $(GEN_CONTRACTS_OUTPUT_PATH)/slash_indicator/slash_indicator.go

ronin:
	CGO_CFLAGS="-O -D__BLST_PORTABLE__" $(GORUN) build/ci.go install ./cmd/ronin
	@echo "Done building."
	@echo "Run \"$(GOBIN)/ronin\" to launch ronin."

bootnode:
	$(GORUN) build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch bootnode."

all:
	$(GORUN) build/ci.go install

android:
	$(GORUN) build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/ronin.aar\" to use the library."
	@echo "Import \"$(GOBIN)/ronin-sources.jar\" to add javadocs"
	@echo "For more info see https://stackoverflow.com/questions/20994336/android-studio-how-to-attach-javadoc"

ios:
	$(GORUN) build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Geth.framework\" to use the library."

test: all
	$(GORUN) build/ci.go test

lint: ## Run linters.
	$(GORUN) build/ci.go lint

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go install golang.org/x/tools/cmd/stringer@latest
	env GOBIN= go install github.com/kevinburke/go-bindata/go-bindata@latest
	env GOBIN= go install github.com/fjl/gencodec@latest
	env GOBIN= go install github.com/golang/protobuf/protoc-gen-go@latest
	env GOBIN= go install ./cmd/abigen
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

ronin-cross: ronin-linux geth-darwin geth-windows geth-android geth-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/ronin-*

ronin-linux: ronin-linux-386 geth-linux-amd64 geth-linux-arm geth-linux-mips64 geth-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/ronin-linux-*

ronin-linux-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/ronin
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/ronin-linux-* | grep 386

ronin-linux-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/ronin
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/ronin-linux-* | grep amd64

ronin-linux-arm: ronin-linux-arm-5 geth-linux-arm-6 geth-linux-arm-7 geth-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/ronin-linux-* | grep arm

ronin-linux-arm-5:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/ronin
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/ronin-linux-* | grep arm-5

ronin-linux-arm-6:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/ronin
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/ronin-linux-* | grep arm-6

ronin-linux-arm-7:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/ronin
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/ronin-linux-* | grep arm-7

ronin-linux-arm64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/ronin
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/ronin-linux-* | grep arm64

ronin-linux-mips:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/ronin
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/ronin-linux-* | grep mips

ronin-linux-mipsle:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/ronin
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/ronin-linux-* | grep mipsle

ronin-linux-mips64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/ronin
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/ronin-linux-* | grep mips64

ronin-linux-mips64le:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/ronin
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/ronin-linux-* | grep mips64le

ronin-darwin: ronin-darwin-386 geth-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/ronin-darwin-*

ronin-darwin-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/ronin
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/ronin-darwin-* | grep 386

ronin-darwin-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/ronin
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/ronin-darwin-* | grep amd64

ronin-windows: ronin-windows-386 geth-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/ronin-windows-*

ronin-windows-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/ronin
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/ronin-windows-* | grep 386

ronin-windows-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/ronin
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/ronin-windows-* | grep amd64
