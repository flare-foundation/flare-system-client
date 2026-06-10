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

RUN apt-get -y update && apt-get -y install curl

WORKDIR /app
COPY --from=builder /app/flare_tcp .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./flare_tcp", "--config", "config.toml" ]
