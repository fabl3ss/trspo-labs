FROM golang:latest

LABEL mainteiner="server"

WORKDIR /server

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

RUN find . -name "*.go" -type f -delete

CMD ["./main"]