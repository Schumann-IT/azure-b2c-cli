.PHONY: test clean build

ARCH := $(shell uname -m)
ifeq ($(ARCH),x86_64)
	ARCH := amd64
endif
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
VERSION := $(shell curl -s https://api.github.com/repos/Schumann-IT/go-ieftool/releases/latest | grep "tag_name" | awk '{print $$2}' | sed 's|[\"\,]*||g')

clean:
	@rm -Rf build
	@rm -f ./ieftool

build:
	@go build -o azb2c

install: ieftool
	@sudo mv ieftool /usr/local/bin/ieftool

ieftool:
	@curl -s -L -o ieftool https://github.com/Schumann-IT/go-ieftool/releases/download/$(VERSION)/ieftool-$(OS)-$(ARCH)
	@chmod +x ieftool
	@if [ "$(OS)" = "darwin" ]; then\
        xattr -d com.apple.quarantine ./ieftool /dev/null 2>&1 | true; \
    fi

check:
	@if [[ "" == "$(GPG_FINGERPRINT)" ]]; then echo "please provide GPG_FINGERPRINT"; exit 1; fi
	@if [[ "" == "$(GITHUB_TOKEN)" ]]; then echo "please provide GITHUB_TOKEN"; exit 1; fi

release: check
	@goreleaser release --clean --timeout 2h --verbose --parallelism 4 --skip=publish --snapshot