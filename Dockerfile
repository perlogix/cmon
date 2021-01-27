FROM golang

ARG GOOS

RUN mkdir -p /go/src/github.com/yeticloud/yeti-discover
RUN go get -u golang.org/x/lint/golint
RUN curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.6.1
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.34.1
WORKDIR /go/src/github.com/yeticloud/yeti-discover

COPY ./ ./

RUN ls -la ./

RUN make lint
RUN make GOOS=${GOOS} build
