FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copiar módulos y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o notification-service ./cmd/api

# Imagen final
FROM alpine:3.18

WORKDIR /app

# Copiar el binario compilado y el archivo .env
COPY --from=builder /app/notification-service .
COPY --from=builder /app/.env .

# Exponer el puerto que usará la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./notification-service"]