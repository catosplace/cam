BINARY_NAME=cam
BIN="./bin"
SRC=$(shell find . -name "*.go")

default: install

fmt:
	@echo "Format Checking"
	@gofmt -w $(SRC)
	@echo "[OK] Formating Checked!"

generate:
	@go generate ./... 
	@echo "[OK] Files added to embed box!"

lint:
	@golangci-lint run

security:
	@gosec ./...
	@echo "[OK] Go security check was completed!"

build: fmt lint generate security
	@go build -o $(BIN)/$(BINARY_NAME) -v 
	@echo "[OK] App binary was created!"

install: generate
	@go install
	@echo "[OK] App binary was installed!"

clean:
	@rm -rf $(BIN)
	@echo "[OK] App binary was removed!"

PHONY: fmt generate security build install clean
