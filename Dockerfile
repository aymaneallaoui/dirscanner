FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o dirscanner main.go

FROM alpine:3.14
WORKDIR /root/
COPY --from=builder /app/dirscanner .
ENTRYPOINT ["./dirscanner"]
