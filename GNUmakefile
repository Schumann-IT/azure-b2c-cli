.PHONY: test clean build

ARCH := $(shell uname -m)
ifeq ($(ARCH),x86_64)
	ARCH := amd64
endif
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
TAG := $(shell curl -s https://api.github.com/repos/Schumann-IT/azure-b2c-cli/releases/latest | grep "tag_name" | awk '{print $$2}' | sed 's|[\"\,]*||g')
VERSION := $(shell  echo $(TAG) | tr -d v)

clean:
	@rm -Rf build
	@rm -Rf dist
	@rm -f ./azure-b2c-cli*

build:
	@go build -o azure-b2c-cli

install: azure-b2c-cli
	@sudo mv azure-b2c-cli /usr/local/bin/azure-b2c-cli

azure-b2c-cli:
	@curl -s -L -o azure-b2c-cli_$(TAG).zip https://github.com/Schumann-IT/azure-b2c-cli/releases/download/$(TAG)/azure-b2c-cli_$(VERSION)_$(OS)_$(ARCH).zip
	@unzip azure-b2c-cli_$(TAG).zip
	@mv azure-b2c-cli_$(TAG) azure-b2c-cli
	@if [ "$(OS)" = "darwin" ]; then\
        xattr -d com.apple.quarantine ./azure-b2c-cli /dev/null 2>&1 | true; \
    fi

check:
	@if [[ "" == "$(GPG_FINGERPRINT)" ]]; then echo "please provide GPG_FINGERPRINT"; exit 1; fi
	@if [[ "" == "$(GITHUB_TOKEN)" ]]; then echo "please provide GITHUB_TOKEN"; exit 1; fi

release: check
	@goreleaser release --clean --timeout 2h --verbose --parallelism 4