FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go mod download && go build -o it-revolution-test .

FROM alpine

WORKDIR /app

COPY --from=builder /build/it-revolution-test /app/it-revolution-test

CMD ["./it-revolution-test"]
