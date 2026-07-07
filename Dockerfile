##########################
# Stage 1
##########################

FROM golang:1.24-alpine AS builder

WORKDIR /src

#COPY go.mod .
#COPY go.sum .
COPY go.mod ./
RUN go mod download

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build \
        -trimpath \
        -ldflags="-s -w" \
        -o /kea-ui \
        ./cmd/server

##########################
# Stage 2
##########################

FROM alpine:3.22

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /kea-ui .
COPY configs/config.yaml ./configs/config.yaml

EXPOSE 8080

ENTRYPOINT ["/app/kea-ui"]
