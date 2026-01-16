fmt:
	go fmt ./...

install:
	go install github.com/meetwithabhishek/rr

# Install with github.com/charmbracelet/fang support. github.com/charmbracelet/fang makes the CLI app fancy
# with more vibrant output, but kind of makes the CLI a bit slower.	
install-fang:
	go install -tags fang github.com/meetwithabhishek/rr

build:
	go build . 

# Build with github.com/charmbracelet/fang support. github.com/charmbracelet/fang makes the CLI app fancy
# with more vibrant output, but kind of makes the CLI a bit slower.	
build-fang:
	go build -tags fang .

lint:
	golangci-lint run --timeout 5m

tidy:
	go mod tidy

vendor: tidy
	go mod vendor

clean:
	rm -rf rr

.PHONY: fmt lint install install-fang build build-fang tidy vendor clean
