MAIN_PACKAGE := yeti-discover
BUILT_ON := $(shell date)
GO_LINUX := GOOS=linux GOARCH=amd64
GO_OSX := GOOS=darwin GOARCH=amd64
LDFLAGS := '-s -w -X "main.builtOn=$(BUILT_ON)"'

osx:
	CGO_ENABLED=0 $(GO_OSX) go build -a -installsuffix cgo -o $(MAIN_PACKAGE) -ldflags $(LDFLAGS) .

linux:
	CGO_ENABLED=0 $(GO_LINUX) go build -a -installsuffix cgo -o $(MAIN_PACKAGE) -ldflags $(LDFLAGS) .

clean:
	find . -name *_gen.go -type f -exec rm {} \;
	rm -f ./$(MAIN_PACKAGE)

lint: 
	${GOPATH}/bin/golint network data db packages containers cloud config system
	${GOPATH}/bin/golangci-lint run

run:
	go run main.go

fmt:
	go fmt ./...
