
.PHONY: clean test help default

M=$(shell printf "\033[34;1m▶\033[0m")

default: get-deps build

help:
	@echo 'Management commands for smartapp-go:'
	@echo
	@echo 'Usage:'
	@echo '    make get-deps        runs glide install, mostly used for ci.'
	@echo '    make fmt             runs go fmt.'
	@echo '    make test            Run tests on a compiled project.'
	@echo '    make generate        Generates go code for working with SmartThings.'
	@echo

get-deps: ; $(info $(M) Installing project dependencies…) @
	glide install

fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		gofmt -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

test: ; $(info $(M) Running tests with GOPATH=${GOPATH}…) @
	go test $$(glide nv)

generate: ; $(info $(M) Generate SmartThings API bindings with GOPATH=${GOPATH}…) @
	third_party/swagger/0.16.0/swagger_darwin_386 generate client -f=third_party/st-api.yml -t=./pkg/smartthings --skip-validation --api-package=operations --client-package=client
	third_party/swagger/0.16.0/swagger_darwin_386 generate model --spec=third_party/smartapps-v1.yml --model-package=smartapp --target=./pkg --skip-validation
