FROM golang:1.23-alpine as builder

WORKDIR /app

RUN apk add --no-cache git curl && \
    go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air"]
