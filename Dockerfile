FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o main cmd/ordersystem/main.go

EXPOSE 8000
EXPOSE 50051
EXPOSE 8080

CMD ["./main"]
