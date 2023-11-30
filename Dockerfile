# Use official Golang image from DockerHub
FROM golang:1.21.4
ENV GOPROXY=https://goproxy.io,direct
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download


COPY . .
RUN go install github.com/cespare/reflex@latest

EXPOSE 5050

# Makefile as entrypoint
ENTRYPOINT ["make"]
CMD ["dev"]