# Build the Go application
FROM golang:1.24-alpine AS go-builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Copy go mod files first (for better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy all Go source code
COPY . .

# Copy the built frontend from the previous stage
COPY --from=web-builder /build/server/router/frontend/dist ./server/router/frontend/dist

# Build the Go application
# CGO_ENABLED=0 for a static binary
# -ldflags="-w -s" to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -X github.com/imrany/spindle/internal/version.version=$(git describe --tags --always --dirty)" \
    -o spindle \
    ./main.go

# Stage 3: Final runtime image
FROM alpine:3.19

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata && \
    addgroup -S spindle && \
    adduser -S spindle -G spindle

WORKDIR /opt/spindle

# Copy the binary from builder stage
COPY --from=go-builder /app/spindle .

# Switch to non-root user
USER spindle

# Expose the port (adjust if different)
EXPOSE 5230

# Set environment variables
ENV ADDR=0.0.0.0
ENV PORT=5230

LABEL org.opencontainers.image.description DESCRIPTION="spindle is an open-source, lightweight web crawler and scraper. It can discover links on the web (crawler) and extract structured data from webpages (scraper)."

# Start the application
ENTRYPOINT ["./spindle"]