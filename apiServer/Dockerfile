FROM golang:1.16.7-alpine3.14

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]