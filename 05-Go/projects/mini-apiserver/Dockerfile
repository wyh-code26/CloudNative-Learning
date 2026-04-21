FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod main.go ./
RUN go build -o mini-apiserver main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/mini-apiserver .
EXPOSE 8081
CMD ["./mini-apiserver"]
