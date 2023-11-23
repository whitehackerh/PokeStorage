FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./internal ./internal
COPY ./cmd ./cmd

RUN go build -o main ./cmd/server/main.go

EXPOSE 8080

CMD ["./main"]
