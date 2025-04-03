# Build stage
FROM golang:1.24-alpine AS base

# Instalar dependencias necesarias para CGO y SQLite
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy go mod and sum files
COPY go.* ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 go build -o app ./cmd/server

# Run stage
FROM alpine:latest

# Instalar SQLite
RUN apk add --no-cache sqlite-libs

WORKDIR /app

# Copy the binary from builder
COPY --from=base /app/app /app/app

# Copy db migrations
COPY db/ /app/db/

# Expose API port (adjust if needed)
EXPOSE 8080
ENTRYPOINT /app/app