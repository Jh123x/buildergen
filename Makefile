.PHONY: install
install:
	go mod tidy

.PHONY: test
test:
	go test -coverprofile cov.out -v -cover ./internal/consts/... ./internal/generation/... ./internal/utils/... ./internal/parser ./internal/writer ./internal/cmd/...

.PHONY: coverage
coverage: test
	go tool cover -func=cov.out
