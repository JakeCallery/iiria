FROM golang:1.16.7 AS builder
WORKDIR /go/src/github.com/jakecallery/iiria/apiServer
COPY . .
RUN CGO_ENABLED=0 go build -o apiserver_app

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/jakecallery/iiria/apiServer/apiserver_app ./
CMD ["./apiserver_app"]
