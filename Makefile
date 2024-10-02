.PHONY: all clean

BINARY_NAME=unite-workload-example

all: build-amd64 build-arm64

build-amd64:
	GOARCH=amd64 go build -o $(BINARY_NAME)-amd64 cmd/main.go

build-arm64:
	GOARCH=arm64 go build -o $(BINARY_NAME)-arm64 cmd/main.go

clean:
	rm -f $(BINARY_NAME)-amd64 $(BINARY_NAME)-arm64