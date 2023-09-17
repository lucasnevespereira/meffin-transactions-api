FROM golang:1.20 AS builder

WORKDIR /app
ADD . /app

RUN go build -o meffin-transactions-api cmd/server/main.go

FROM ubuntu:latest AS launcher
COPY --from=builder /app .
CMD ["./meffin-transactions-api"]