.DEFAULT_GOAL := fmt

update:
	@echo "Going into sources..."
    cd src/
	@echo "Updating dependencies..."
	go get -u
	@echo "Cleaning up..."
	go mod tidy

fmt:
	gofmt -s -w .
