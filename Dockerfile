# build executable
FROM golang:1.25.1-trixie@sha256:61226c61f37cb86253c4ac486ef22c47f14bfddb8f60bb4805bfc165001be758 AS builder

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
