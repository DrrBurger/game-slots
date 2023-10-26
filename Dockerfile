FROM golang:1.21

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o game-slots cmd/server/main.go

EXPOSE 8080

CMD ["./game-slots"]