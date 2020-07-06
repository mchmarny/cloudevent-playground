.PHONY: mod test run lint tag clean
all: test

mod: ## Updates the go modules and vendors all dependencies 
	go mod tidy

test: mod ## Tests the entire project 
	go test -count=1 -race ./...

run: mod ## Runs the uncompiled code
	go run rand.go event.go main.go 

lint: ## Lints the entire project 
	golangci-lint run --timeout=3m

clean: ## Cleans up generated files 
	go clean
	rm -fr ./bin

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk \
		'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
