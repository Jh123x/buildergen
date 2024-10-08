.PHONY: install
install:
	go mod tidy

.PHONY: test
test:
	go test -coverprofile cov.out -v ./internal/cmd/... ./internal/consts/... ./internal/generation/... ./internal/utils/... ./internal/parser

.PHONY: coverage
coverage: test
	go tool cover -func=cov.out
