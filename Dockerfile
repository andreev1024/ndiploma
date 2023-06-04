FROM golang:1.20 as base

# dev stage start

FROM base as dev

# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin#

RUN go install github.com/cosmtrek/air@latest

WORKDIR /opt/app/api
CMD ["air"]

# built stage start

FROM base as built

WORKDIR /go/app/api
COPY . .

ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go build -o /tmp/api-server ./*.go

# busybox stage start

FROM busybox

COPY --from=built /tmp/api-server /usr/bin/api-server
CMD ["api-server", "start"]
