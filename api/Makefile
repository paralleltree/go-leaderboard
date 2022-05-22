define USAGE
--- Available commands ---
help:    Print this help
mockgen: Regenerate mocks defined in internal/contract with go generate
endef

export USAGE

.PHONY: help
help:
	@echo "$${USAGE}"

.PHONY: mockgen
mockgen:
	@go install github.com/golang/mock/mockgen@73266f9366fcf2
	@rm -rf mock/
	@go generate ./...
