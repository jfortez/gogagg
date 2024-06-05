
FROM golang:1.22.2-alpine3.19 AS BUILDER

ENV ADDRESS="127.0.0.1:8000"
ENV CGO_ENABLED=1

RUN apk add --no-cache git \
  # Important: required for go-sqlite3
  gcc \
  # Required for Alpine
  musl-dev


WORKDIR /app

COPY . .
RUN go get -d -v ./...
RUN go build -o app -v


ENTRYPOINT ["./app"]