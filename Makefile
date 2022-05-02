MAIN_PACKAGE := cmon
BUILT_ON := $(shell date)
GOOS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
COMMIT:=$(shell git log -n 1 --pretty=format:"%H" 2>/dev/null | head -c 10)
PACKAGES:=$(shell go list ./... | grep -v /vendor/)
GO_LINUX := GOOS=linux GOARCH=amd64
GO_ARM := GOOS=linux GOARCH=arm64
GO_OSX := GOOS=darwin GOARCH=amd64
GO_WIN := GOOS=windows GOARCH=amd64
VER := 1.1
LDFLAGS := '-s -w -X "github.com/perlogix/cmon/config.builtOn=$(BUILT_ON)" -X "github.com/perlogix/cmon/config.commitHash=$(COMMIT)" -X "github.com/perlogix/cmon/config.version=$(VER)"'

default: build

build:
	GOOS=$(GOOS) CGO_ENABLED=0 go build -a -installsuffix cgo -o $(MAIN_PACKAGE) -ldflags $(LDFLAGS) .

osx:
	CGO_ENABLED=0 $(GO_OSX) go build -a -installsuffix cgo -o $(MAIN_PACKAGE) -ldflags $(LDFLAGS) .

linux:
	CGO_ENABLED=0 $(GO_LINUX) go build -a -installsuffix cgo -o $(MAIN_PACKAGE) -ldflags $(LDFLAGS) .

arm:
	CGO_ENABLED=0 $(GO_ARM) go build -a -installsuffix cgo -o $(MAIN_PACKAGE) -ldflags $(LDFLAGS) .

windows:
	CGO_ENABLED=0 $(GO_WIN) go build -a -installsuffix cgo -o $(MAIN_PACKAGE).exe -ldflags $(LDFLAGS) .

release:
	upx $(MAIN_PACKAGE) || true
	upx $(MAIN_PACKAGE).exe || true

clean:
	find . -name *_gen.go -type f -delete
	rm -f ./$(MAIN_PACKAGE)
	rm -f ./*.rpm
	rm -f ./*.deb

gofmt:
	go fmt ./...

lint: gofmt
	$(GOPATH)/bin/staticcheck $(PACKAGES)
	$(GOPATH)/bin/golangci-lint run
	$(GOPATH)/bin/gosec -quiet -no-fail ./...

run:
	go run main.go

update-deps:
	go get -u ./...
	go mod tidy

docker:
	sudo docker rm -f $(MAIN_PACKAGE)-build 2>/dev/null
	sudo docker build --target build --build-arg GOOS=$(GOOS) -t $(MAIN_PACKAGE)-build .
	sudo docker create --name $(MAIN_PACKAGE)-build $(MAIN_PACKAGE)-build
	sudo docker cp $(MAIN_PACKAGE)-build:/go/src/build/$(MAIN_PACKAGE) ./
	sudo docker rm -f $(MAIN_PACKAGE)-build

pkgs:
	sudo docker rm -f $(MAIN_PACKAGE)-pkgs 2>/dev/null
	sudo docker build --build-arg VER=$(VER) -t $(MAIN_PACKAGE)-pkgs -f Dockerfile-pkgs .
	sudo docker create --name $(MAIN_PACKAGE)-pkgs $(MAIN_PACKAGE)-pkgs
	sudo docker cp $(MAIN_PACKAGE)-pkgs:/packaging/$(MAIN_PACKAGE)-$(VER)-amd64.rpm ./
	sudo docker cp $(MAIN_PACKAGE)-pkgs:/packaging/$(MAIN_PACKAGE)-$(VER)-amd64.deb ./
	sudo docker rm -f $(MAIN_PACKAGE)-pkgs
