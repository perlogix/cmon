MAIN_PACKAGE := yeti-discover
BUILT_ON := $(shell date)
GOOS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
GO_LINUX := GOOS=linux GOARCH=amd64
GO_OSX := GOOS=darwin GOARCH=amd64
GO_WIN := GOOS=darwin GOARCH=amd64
LDFLAGS := '-s -w -X "main.builtOn=$(BUILT_ON)"'

build:
	GOOS=$(GOOS) CGO_ENABLED=0 go build -a -installsuffix cgo -o $(MAIN_PACKAGE) -ldflags $(LDFLAGS) .

osx:
	CGO_ENABLED=0 $(GO_OSX) go build -a -installsuffix cgo -o $(MAIN_PACKAGE) -ldflags $(LDFLAGS) .

linux:
	CGO_ENABLED=0 $(GO_LINUX) go build -a -installsuffix cgo -o $(MAIN_PACKAGE) -ldflags $(LDFLAGS) .

windows:
	CGO_ENABLED=0 $(GO_WIN) go build -a -installsuffix cgo -o $(MAIN_PACKAGE).exe -ldflags $(LDFLAGS) .

clean:
	find . -name *_gen.go -type f -exec rm {} \;
	rm -f ./$(MAIN_PACKAGE)

lint:
	fmt
	#${GOPATH}/bin/golint network data db packages containers cloud config system security
	${GOPATH}/bin/golint -set_exit_status ./...
	${GOPATH}/bin/golangci-lint run

run:
	go run main.go

fmt:
	go fmt ./...

docker:
	sudo docker build --build-arg GOOS=$(GOOS) -t yeti-discover-build .
	sudo docker create --name yeti-discover-build yeti-discover-build
	sudo docker cp yeti-discover-build:/go/src/github.com/yeticloud/yeti-discover/yeti-discover ./
	sudo docker rm -f yeti-discover-build
