FROM golang

ENV GOPATH /go
ARG GOOS

RUN mkdir -p /go/src/github.com/yeticloud/yeti-discover
RUN go get -u golang.org/x/lint/golint
RUN go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

WORKDIR /go/src/gitlab.com/yeticloud/yeti-discover

COPY ./* ./

RUN make lint
RUN make GOOS=${GOOS} build