# Stage 1: Build
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Disable CGO, strip debug symbols for a smaller, faster binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/server

# Stage 2: Final — use distroless instead of alpine for a leaner image
FROM gcr.io/distroless/static-debian12
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]