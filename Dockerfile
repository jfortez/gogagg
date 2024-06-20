# build stage
FROM golang:1.22.2-alpine3.19 AS BUILDER

ENV CGO_ENABLED=1

RUN apk add --no-cache git \
  gcc \
  musl-dev \
  upx

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
RUN go build -o app -v -ldflags="-s -w"
RUN upx app

# node stage
FROM node:20.9.0-alpine3.19 AS NODE

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build  # This should include the Tailwind build script

# postgres stage
FROM postgres:16.3-alpine3.19 AS POSTGRES

ENV POSTGRES_PASSWORD=root
ENV POSTGRES_USER=root
ENV POSTGRES_DB=gogag

COPY ./db/init.sql /docker-entrypoint-initdb.d/db.sql

EXPOSE 5432

# final stage
FROM alpine:3.19

LABEL NAME="gogagg"

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=BUILDER /app/app .
COPY --from=POSTGRES /docker-entrypoint-initdb.d/db.sql /docker-entrypoint-initdb.d/db.sql

ENTRYPOINT ["./app"]
