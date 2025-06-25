# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy source code
COPY backend/ .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /

# Copy the binary from builder
COPY --from=builder /main .
COPY --from=builder /app/migrations /migrations

EXPOSE 8080

CMD ["/main"]