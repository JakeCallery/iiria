FROM golang:1.16.7 AS builder
WORKDIR /go/src/github.com/jakecallery/iiria/worker
COPY . .
RUN CGO_ENABLED=0 go build -o worker_app

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/jakecallery/iiria/worker/worker_app ./
CMD ["./worker_app"]
