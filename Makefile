.DEFAULT_GOAL := help

COMMIT_SHA?=$(shell git rev-parse --short HEAD)
DOCKER?=docker
REGISTRY?=jonascheng
# is Windows_NT on XP, 2000, 7, Vista, 10...
ifeq ($(OS),Windows_NT)
GOOS?=windows
RACE=""
else
GOOS?=$(shell uname -s | awk '{print tolower($0)}')
GORACE="-race"
endif

.PHONY: setup
setup:	## setup go modules
	go mod tidy

.PHONY: clean
clean:	## cleans the binary
	go clean
	rm -rf ./bin
	rm -rf server-v1.*
	rm -rf server-v2.*

.PHONY: run
run: setup server-key ## runs go run the application
	go run ${GORACE} cmd/main.go

.PHONY: server-key
server-key:	## setup server key
	## Key considerations for algorithm RSA ≥ 1024-bit
	if [ ! -f server-v1.key ]; then openssl genrsa -out server-v1.key 1024; fi;
	## Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
	if [ ! -f server-v1.crt ]; then openssl req -new -x509 -key server-v1.key -out server-v1.crt -days 7 -subj "/C=TW/ST=Taipei/L=Test/O=Test/OU=Test/CN=localhost/emailAddress=Test@email"; fi;
	## Key considerations for algorithm RSA ≥ 1024-bit
	if [ ! -f server-v2.key ]; then openssl genrsa -out server-v2.key 1024; fi;
	## Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
	if [ ! -f server-v2.crt ]; then openssl req -new -x509 -key server-v2.key -out server-v2.crt -days 7 -subj "/C=TW/ST=Taipei/L=Test/O=Test/OU=Test/CN=localhost/emailAddress=Test@email"; fi;

.PHONY: help
help: ## prints this help message
	@echo "Usage: \n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
