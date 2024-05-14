FROM golang:1.22 as build
WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download
COPY /config ./config
COPY /domain ./domain
COPY /handler ./handler
COPY /router ./router
COPY main.go ./
RUN GOOS=linux go build -a -o main main.go
