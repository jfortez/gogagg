# Etapa de construcción
FROM golang:1.22.2-alpine AS BUILDER

ENV CGO_ENABLED=1

RUN apk add --no-cache git \
  gcc \
  musl-dev \
  upx

WORKDIR /app

# Copiar los archivos del proyecto

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
RUN go build -o app -v -ldflags="-s -w"
RUN upx app
# Etapa de producción
FROM alpine:latest
LABEL NAME="gogagg"
RUN apk update && apk add --no-cache ca-certificates
WORKDIR /app

# Copiar el binario compilado desde la etapa de construcción
COPY --from=BUILDER /app/app .

# Exponer el puerto
EXPOSE ${PORT}

# Definir las variables de entorno
ENV PORT=${PORT}
ENV SECRET_KEY=${SECRET_KEY}
ENV HOST=${HOST}

# Ejecutar la aplicación
CMD ["./app"]
