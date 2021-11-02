FROM golang:1.17.2-alpine3.14 AS build-env
WORKDIR /app

# Copy file
COPY *.go ./

RUN go build -o /output

EXPOSE 8080

cmd ["/output"]