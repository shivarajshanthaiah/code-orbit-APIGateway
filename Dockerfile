FROM golang:1.22.3

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
COPY .env .env

RUN go build -o main

ENV $(grep -v '^#' .env | xargs)

CMD ["./main"]
