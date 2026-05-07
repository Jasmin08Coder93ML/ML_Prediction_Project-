# Шаг 1: Сборка
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Шаг 2: Запуск
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# Копируем .env файл, если он у вас есть
COPY .env . 

EXPOSE 8080
CMD ["./main"]
