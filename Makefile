
GOTEST=go test
BINARY_NAME=goframe

.PHONY: test build run clean go-lint

build: ## Build project and put the output binary in bin/
	mkdir -p bin/
	GO111MODULE=on go build -o bin/$(BINARY_NAME) .

clean: ## Remove build related file
	rm -fr ./bin

run: build
	./bin/$(BINARY_NAME)

## Test:
test: 
	$(GOTEST) -cover -v -race ./... 

go-lint:
	docker run -t --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint run -v