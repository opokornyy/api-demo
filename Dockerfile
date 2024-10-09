FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go build \
    -o ./bin/api-demo \
    ./cmd/api-demo

FROM alpine as production

WORKDIR /app

COPY --from=builder /app/bin/api-demo /app

ENTRYPOINT [ "/app/api-demo" ]
