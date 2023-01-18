AUTH_APP=authApp

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

PACKAGES = $(shell go list ./...)

fmt:
	@echo "Executing go fmt"
	go fmt $(PACKAGES)
	gofumpt -w .

build:
	@echo "Building auth binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_APP} ./cmd
	@echo "Done!"

remove_binary:
	@echo "Removing binary file"
	rm ${AUTH_APP} || true
