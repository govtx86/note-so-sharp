# backend/Dockerfile

# --- Build Stage ---
FROM golang:1.26.5-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy backend source code
COPY . .

# Build lightweight binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# --- Final Deployment Stage ---
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]