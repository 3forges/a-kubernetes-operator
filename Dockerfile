# ---- Build stage ----
FROM debian:stable-20260223-slim AS builder

RUN mkdir -p /operator
WORKDIR /operator

# Download dependencies first (better caching)
COPY operator ./
# RUN go mod download
# RUN go mod tidy
RUN ls -alh /operator/operator

ENTRYPOINT [ "/operator/operator" ]
# ENTRYPOINT ["/operator"]
# ENTRYPOINT ["/bin/sh"]