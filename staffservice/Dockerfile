FROM golang:1.17.2-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /app/src/internal

RUN go build -o ./out/go-example .

EXPOSE 8080

CMD ["./out/go-example"]
