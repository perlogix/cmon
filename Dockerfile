FROM golang

ARG GOOS

RUN mkdir -p /go/src/github.com/yeticloud/yeti-discover
RUN go get -u golang.org/x/lint/golint
#RUN go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0
WORKDIR /go/src/github.com/yeticloud/yeti-discover

COPY ./* ./

RUN ls -la ./

RUN make lint
RUN make GOOS=${GOOS} build
