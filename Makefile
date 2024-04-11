fmt:
	go fmt
:PHONY: fmt

lint:
	fmt
	staticcheck
:PHONY: lint

vet: fmt
		go vet ./...
			shadow ./...
.PHONY: vet

build: vet
		go mod tidy
		go build
.PHONY: build