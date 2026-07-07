FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o kea-ui ./cmd/server

CMD ["/app/kea-ui"]
