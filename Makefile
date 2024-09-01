.PHONY: build
build:
	go build -o bin/mcdctl cmd/cli/main.go

.PHONY: run
run:
	go run cmd/cli/main.go localhost
