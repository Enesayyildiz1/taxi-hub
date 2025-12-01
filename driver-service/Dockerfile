FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN  go mod download

COPY . .

RUN go build -o driver ./cmd/server/main.go

EXPOSE 8081

CMD ["./driver"]