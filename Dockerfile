FROM golang:1.22 as build
WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download
COPY /cmd ./cmd
COPY /config ./config
COPY /internal ./internal
RUN GOOS=linux go build -a -o main cmd/api/main.go
