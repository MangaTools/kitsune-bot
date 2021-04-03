FROM golang:1.16.3-alpine3.13

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o main ./cmd/main.go

CMD ["./main"]