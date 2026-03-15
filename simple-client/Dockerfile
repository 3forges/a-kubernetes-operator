# ---- Build stage ----
    FROM golang:1.22 AS builder

    WORKDIR /operator
    
    # Download dependencies first (better caching)
    COPY go.mod go.sum ./
    # RUN go mod download
    # RUN go mod tidy
    
    # Copy the rest of the source
    COPY . .
    
    # CGO_ENABLED=0 → ensures the binary is fully static and works with distroless.
    # Build a static binary
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
        go build -ldflags="-s -w" -o operator
    
    # ---- Runtime stage (distroless) ----
    FROM gcr.io/distroless/static-debian12
    # FROM gcr.io/distroless/base-debian12 AS RELEASE
    
    WORKDIR /
    
    COPY --from=builder /operator/operator /operator
    
    USER nonroot:nonroot
    
    ENTRYPOINT ["/operator"]