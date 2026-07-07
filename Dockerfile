FROM golang:1.24-alpine AS builder

WORKDIR /src

COPY . .

COPY go.mod ./
RUN go mod download

RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build \
    -o /kea-ui \
    ./cmd/server

FROM alpine:3.22

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /kea-ui .

EXPOSE 8080

ENTRYPOINT ["/app/kea-ui"]
