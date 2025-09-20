# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Cache Go modules
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy source code
COPY . .

# Build the binary
# Inject version from build arg (default: dev)
ARG VERSION=dev
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -X github.com/imrany/spindle/internal/version.version=${VERSION}" \
    -o spindle \
    ./main.go

# Stage 2: Final runtime image
FROM alpine:3.19

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata && \
    addgroup -S spindle && \
    adduser -S spindle -G spindle

WORKDIR /opt/spindle

# Copy binary
COPY --from=builder /app/spindle .

# Run as non-root
USER spindle

# Expose default port
EXPOSE 5020

# Set defaults
ENV ADDR=0.0.0.0 \
    PORT=5020

LABEL org.opencontainers.image.title="Spindle" \
      org.opencontainers.image.description="Spindle is an open-source, lightweight web crawler and scraper. It can discover links (crawler) and extract structured data (scraper)." \
      org.opencontainers.image.source="https://github.com/imrany/spindle"

# Start the app
ENTRYPOINT ["./spindle"]
