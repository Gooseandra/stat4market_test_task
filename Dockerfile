FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM debian:bullseye-slim

COPY --from=builder /app/main /main

EXPOSE 8081

CMD ["/main"]