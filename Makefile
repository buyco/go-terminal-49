DOCKER_BUILD := docker run --rm -u `id -u` -v ${PWD}:/sdk openapitools/openapi-generator-cli:v4.3.1 generate -i sdk/api_files/terminal49_openapi.yml
GO_CLIENT := -g go -o /sdk/terminal49 \
			--git-repo-id=go-terminal49 --git-user-id=buyco \
			--additional-properties=packageName=terminal49 \
			--additional-properties=isGoSubmodule=true \
			--additional-properties=generateInterfaces=true


## generate: Clean and generate SDK from file.
generate:
	${MAKE} clean
	${MAKE} go-sdk

go-sdk:
	${DOCKER_BUILD} ${GO_CLIENT}

clean:
	rm -rf ./terminal49

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo