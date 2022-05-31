# syntax=docker/dockerfile:1

FROM golang:1.16-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/mikromon-core

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
# RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/mikromon-core ./cmd/core/main.go

# Start fresh from a smaller image
FROM alpine:3.16
RUN apk add ca-certificates

COPY --from=build_base /tmp/mikromon-core/out/mikromon-core /app/mikromon-core

# Run the binary program produced by `go install`
CMD ["/app/mikromon-core"]


# docker build -t mortimor88/mikromon-core:latest -t mortimor88/mikromon-core:1.2 .