# Stage 1: Build the application
FROM docker.io/golang:1.26-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/myapp ./cmd

# Stage 2: Run the application
FROM docker.io/alpine:3.22

RUN addgroup -S app && adduser -S app -G app

COPY --from=builder /app/myapp /myapp

EXPOSE 8080

USER app

HEALTHCHECK --interval=10s --timeout=3s --start-period=10s --retries=3 \
    CMD wget -q --spider http://127.0.0.1:8080/healthz || exit 1

ENTRYPOINT ["/myapp"]
