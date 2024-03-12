run:
	@go run main.go

build:
	@mkdir -p bin
	@go build -o bin/monkey main.go
	@echo "Build complete"
	@echo "Run ./bin/monkey to start the program"
	@echo "Or add the bin directory to your PATH"

test:
	@go test -v ./...
