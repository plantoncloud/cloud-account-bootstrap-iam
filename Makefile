build_cmd=go build

.PHONY: deps
deps:
	go mod download

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: build
build: deps vet

.PHONY: test
test:
	go test -race -v -count=1 ./...
