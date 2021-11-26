FROM golang:latest

RUN go install github.com/cespare/reflex@latest

RUN export PATH=$PATH:$(go env GOPATH)/bin

WORKDIR /app

COPY . .

RUN go mod download


CMD reflex -r '\.go$' -s -- sh -c 'go run main.go'
