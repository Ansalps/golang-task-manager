# Build stage
FROM golang:1.22.2-alpine3.19 AS builder

RUN apk update && apk upgrade --no-cache

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary 
RUN go build -o taskmanager ./cmd/main.go

# Final lightweight image
FROM alpine:latest

RUN apk update && apk upgrade --no-cache

WORKDIR /app

COPY --from=builder /app/taskmanager .
COPY .env .

EXPOSE 8080

CMD ["./taskmanager"]