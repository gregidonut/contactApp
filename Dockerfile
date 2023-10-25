FROM golang:latest

WORKDIR /app

RUN git clone https://github.com/gregidonut/contactApp

WORKDIR /app/contactApp

RUN go build -o app cmd/web/main.go

EXPOSE 8080

CMD ["./app"]