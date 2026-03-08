.PHONY: install
install:
	go mod tidy

.PHONY: test
test:
	go test -coverprofile cov.out -v -cover ./... 

.PHONY: coverage
coverage: test
	go tool cover -func=cov.out

.PHONY: benchmark
benchmark:
	go test -bench=. -benchmem ./examples/...
