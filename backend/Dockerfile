FROM golang:1.21-alpine
WORKDIR /app/backend/src
COPY . .
RUN go mod download
CMD ["go", "run", "main.go"]
