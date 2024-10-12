BINARY=todo-cli

build:
	go build -o bin/$(BINARY) cmd/main.go

test:
	go test ./todo -v