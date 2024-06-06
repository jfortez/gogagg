
# build stage
FROM golang:1.22.2-alpine3.19 AS BUILDER

ENV ADDRESS="127.0.0.1:8000"
ENV CGO_ENABLED=1

RUN apk add --no-cache git \
  # Important: required for go-sqlite3
  gcc \
  # Required for Alpine
  musl-dev \
  upx


WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download -x

COPY . .
RUN go build \
    -o app -v \
    -ldflags="-s -w"
RUN upx app



# final stage
FROM alpine:3.19
LABEL NAME="gogagg"

RUN apk update
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=BUILDER /app .

ENTRYPOINT ["./app"]