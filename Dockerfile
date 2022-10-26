# syntax=docker/dockerfile:1
FROM golang:1.19.0-alpine3.16 as build
WORKDIR /go/src/github.com/ljg-cqu/biance
COPY . .
RUN ls -la
RUN go get -d ./...
RUN apk --update add build-base && GOOS=linux go build -a -o biance .

FROM alpine:3.16
WORKDIR /app/
COPY --from=build /go/src/github.com/ljg-cqu/biance/biance ./
CMD ["./biance"]