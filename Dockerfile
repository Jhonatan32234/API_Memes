# Etapa 1: Construcción
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Compilamos el binario desde la ruta donde está tu main.go
RUN go build -o main ./cmd/api/main.go

# Etapa 2: Ejecución
FROM alpine:latest
WORKDIR /app
# Copiamos el binario y el archivo .env
COPY --from=builder /app/main .
COPY --from=builder /app/cmd/api/.env .

EXPOSE 8080
CMD ["./main"]