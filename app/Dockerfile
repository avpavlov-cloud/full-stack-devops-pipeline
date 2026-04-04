# ЭТАП 1: Сборка (Build)
FROM golang:1.26-alpine AS builder
WORKDIR /build
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# ЭТАП 2: Финальный образ (Production)
FROM alpine:latest
WORKDIR /root/
# Копируем только скомпилированный бинарник из первого этапа
COPY --from=builder /build/main .

EXPOSE 8080
CMD ["./main"]
