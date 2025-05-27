FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags="-s -w" -o flats-parser .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /build/flats-parser .

ENTRYPOINT ["./flats-parser"]
