# ==========================
# Stage 1: Builder
# ==========================
FROM alpine:3.18 AS builder

# Install dependencies for Go and building
RUN apk add --no-cache bash git curl tar xz build-base

# Set Go version and environment variables
ENV GO_VERSION=1.25.5 \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    PATH="/usr/local/go/bin:${PATH}"

# Download and install Go
RUN curl -LO https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (Docker cache optimization)
COPY go.mod go.sum ./

# Download all dependencies specified in go.mod/go.sum
RUN go mod download
RUN go mod verify

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o gin-app .

# ==========================
# Stage 2: Runtime
# ==========================
FROM alpine:3.18

# Add CA certificates for HTTPS support
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/gin-app .

# Expose default Gin port
EXPOSE 8080

# Run the application
CMD ["./gin-app"]
