# build executable
FROM golang:1.21 AS builder

WORKDIR /build

# Copy and download dependencies using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . ./

# Build the applications
RUN go build -o /app/flare_tcp ./client/main/client.go

FROM debian:latest AS execution

ARG deployment=flare
ARG type=voting

RUN apt-get -y update && apt-get -y install curl

WORKDIR /app
COPY --from=builder /app/flare_tcp .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./docker/client/config.toml ./config.toml

CMD ["./flare_tcp", "--config", "config.toml" ]
