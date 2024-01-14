FROM golang:1.21.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

VOLUME /logs

RUN go build -o bin ./cmd/fpl-go-bot

CMD ["./bin"]
