# Этап сборки
FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ВАЖНО: Статическая сборка, без glibc
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notification-sender main.go

# Финальный образ — пустой (без glibc, минимальный размер)
FROM scratch

WORKDIR /app

# Копируем собранный бинарник
COPY --from=builder /app/notification-sender .
COPY --from=builder /app/service ./service
COPY --from=builder /app/.env ./.env
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Устанавливаем переменные окружения (опционально)
ENV DB_HOST=mysql \
    DB_PORT=3306 \
    DB_USER=root \
    DB_PASSWORD=root \
    DB_NAME=drs_db

# Запускаем
ENTRYPOINT ["./notification-sender"]
