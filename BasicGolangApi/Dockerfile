FROM golang:1.17.2-alpine3.14 AS build-env
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy file
COPY *.go ./

RUN go build -o /output

EXPOSE 8080

cmd ["/output"]