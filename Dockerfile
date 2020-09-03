FROM golang:1.14-alpine AS baseGo
ENV CGO_ENABLED 0
RUN apk --no-cache add git
RUN go get -u github.com/go-delve/delve/cmd/dlv
WORKDIR $GOPATH/src/github.com/sunsided/hashicorp-raft-example
COPY go.mod .
COPY go.sum .
RUN go mod download -x
COPY . .
RUN go build -gcflags='all=-N -l' -o example cmd/raft-example/main.go

FROM scratch AS release
LABEL stage=release
COPY --from=baseGo /go/src/github.com/sunsided/hashicorp-raft-example/example /
ENTRYPOINT ["./example"]

FROM scratch AS debug
LABEL stage=debug
COPY --from=baseGo /go/bin/dlv /
COPY --from=release /example /
EXPOSE 40000
ENTRYPOINT ["/dlv", "--listen=127.0.0.1:40000", "--headless=true", "--log", "--api-version=2", "exec", "/example"]

FROM scratch AS mock
LABEL stage=mock

FROM golang:alpine AS dev-env
LABEL stage=dev-env
COPY --from=baseGo /go /go
WORKDIR $GOPATH/src/github.com/sunsided/hashicorp-raft-example/
ENTRYPOINT ["./example"]

FROM release
