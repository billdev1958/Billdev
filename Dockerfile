FROM golang:1.21.0-alpine3.18 
WORKDIR /app
COPY . .

RUN go build -o main main.go

COPY .env .

EXPOSE 8080
CMD ["/app/main"]
