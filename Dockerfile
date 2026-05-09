# Stage 1: Build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# Exponer el puerto que definimos en el código
EXPOSE 8080
CMD ["./main"]