FROM golang

ARG GOOS

RUN mkdir -p /go/src/github.com/perlogix/cmon
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go get -u github.com/securego/gosec/v2/cmd/gosec
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.41.1
WORKDIR /go/src/github.com/perlogix/cmon

COPY ./ ./

RUN make lint
RUN make GOOS=${GOOS} build
