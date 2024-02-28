# Define targets for each service
.PHONY: gen fmt test

all: fmt test

fmt: 
	go fmt ./...
	gofumpt -l -w .

lint: 
	golangci-lint run ./...

test: 
	set -o pipefail && go test -v ./... 

# New gen target to incorporate the build process of different services
gen:
	for service in social-service user-service financial-service accounting-service workspace-service; do \
		cd api/$$service && make && cd ../..; \
	done
	make all
	./generate.sh


precommit: gen fmt test
