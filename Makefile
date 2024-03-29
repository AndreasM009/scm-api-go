GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOMOD_DOWNLOAD=$(GOCMD) mod download
BINARY_PATH=./bin/
BINARY_NAME_WINDOWS=$(BINARY_PATH)api.exe
BINARY_NAME_LINUX=$(BINARY_PATH)api
PROJECT_PATH=./cmd/api

build-windows:
	$(GOBUILD) -a -installsuffix cgo -o $(BINARY_NAME_WINDOWS) $(PROJECT_PATH)

build-linux:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o $(BINARY_NAME_LINUX) $(PROJECT_PATH)

deps:
	$(GOMOD_DOWNLOAD)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME_WINDOWS)
	rm -f $(BINARY_NAME_LINUX)