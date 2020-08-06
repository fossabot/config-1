TOOLS_MOD_DIR = ./tools
TOOLS_DIR = $(abspath ./.tools)

GO_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}\
{{end}}' ./...)

GO_TEST_MIN = go test -v -timeout 30s
GO_TEST = $(GO_TEST_MIN) -race -failfast
GO_TEST_WITH_COVERAGE = $(GO_TEST) -coverprofile=coverage.txt -covermode=atomic

.DEFAULT_GOAL = precommit

$(TOOLS_DIR)/golangci-lint: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: precommit
precommit: test lint ## run steps to make sure code pass on CI

.PHONY: test
test: ## run the test cases
	$(GO_TEST) ./...

.PHONY: test-with-coverage
test-with-coverage: ## run testing with code coverage and output html file
	$(GO_TEST_WITH_COVERAGE) ./... && go tool cover -html=coverage.txt -o coverage.html

.PHONY: lint
lint: $(TOOLS_DIR)/golangci-lint ## lint the source code using golangci
	$(TOOLS_DIR)/golangci-lint run --fix && $(TOOLS_DIR)/golangci-lint run

.PHONY: ci
ci: lint test-with-coverage ## list of task that going to be executed on ci

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
