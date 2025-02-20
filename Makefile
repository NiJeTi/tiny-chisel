GOLANGCI_LINT_IMAGE=golangci/golangci-lint:v1.63-alpine

.PHONY: deps
deps:
	docker pull $(GOLANGCI_LINT_IMAGE)

.PHONY: lint
lint:
	$(MAKE) deps

	docker run -t --rm -v $(PWD):/src -w /src $(GOLANGCI_LINT_IMAGE) golangci-lint run
