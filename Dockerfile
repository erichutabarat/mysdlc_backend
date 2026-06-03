# Stage 1: Build
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/server

# Stage 2: Final
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# Expose the port your Go app runs on
EXPOSE 8000
CMD ["./main"]