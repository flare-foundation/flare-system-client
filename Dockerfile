# checkov:skip=CKV_DOCKER_2: Healthcheck is handled by the container orchestrator
# build executable
FROM golang:1.26.4-trixie@sha256:0dcba0d95dbfb072e9917a106b9e07d7cc298097dc83e9307056ef1889de654d AS builder

WORKDIR /build

# Copy and download dependencies using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . ./

# Build the applications
RUN go build -o /app/flare_tcp ./client/main/client.go

FROM debian:trixie@sha256:72547dd722cd005a8c2aa2079af9ca0ee93aad8e589689135feaed60b0a8c08d AS execution

ARG deployment=flare
ARG type=voting

#hadolint ignore=DL3008
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    && rm -rf /var/lib/apt/lists/*

RUN groupadd -g 10001 app && useradd -u 10001 -g app -m app

WORKDIR /app
COPY --from=builder /app/flare_tcp .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER app

CMD ["./flare_tcp", "--config", "config.toml" ]
