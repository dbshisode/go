#build stage
FROM golang:ubuntu AS builder
RUN apt add --no-cache git
WORKDIR /usr/local/src
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=src Version=0.0.1
EXPOSE 3000
