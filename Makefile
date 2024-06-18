build:
	@go build -o bin/dsh ./cmd/goshell/

run: build
	@./bin/dsh
