FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main ./cmd/main.go

CMD ["./main"]

EXPOSE 8080
