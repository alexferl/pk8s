.PHONY: dev gen run test cover cover-html fmt pre-commit

.DEFAULT: help
help:
	@echo "make dev"
	@echo "	setup development environment"
	@echo "make gen"
	@echo "	generate Kubernetes types"
	@echo "make run"
	@echo "	run app"
	@echo "make test"
	@echo "	run go test"
	@echo "make cover"
	@echo "	run go test with -cover"
	@echo "make cover-html"
	@echo "	run go test with -cover and show HTML"
	@echo "make tidy"
	@echo "	run go mod tidy"
	@echo "make fmt"
	@echo "	run gofumpt"
	@echo "make pre-commit"
	@echo "	run pre-commit hooks"

check-gofumpt:
ifeq (, $(shell which gofumpt))
	$(error "gofumpt not in $(PATH), gofumpt (https://pkg.go.dev/mvdan.cc/gofumpt) is required")
endif

check-pre-commit:
ifeq (, $(shell which pre-commit))
	$(error "pre-commit not in $(PATH), pre-commit (https://pre-commit.com) is required")
endif

dev: check-pre-commit
	pre-commit install

version ?=

gen:
	./hack/gen.sh $(version)

run:
	go build -o cli-bin ./cmd/cli && ./cli-bin

build:
	go build -o cli-bin ./cmd/cli

test:
	go test -v $(shell go list ./... | grep -Ev '/hack|/examples')

cover:
	go test -v -cover $(shell go list ./... | grep -Ev '/hack|/examples')

cover-html:
	go test -v -coverprofile=coverage.out $(shell go list ./... | grep -Ev '/hack|/examples')
	go tool cover -html=coverage.out

tidy:
	go mod tidy

fmt: check-gofumpt
	gofumpt -l -w .

pre-commit: check-pre-commit
	pre-commit
