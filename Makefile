GO=go

all: build
build:
	GOOS=linux GOARCH=arm GOARM=6 $(GO) build -v
	
build-local:
	$(GO) build