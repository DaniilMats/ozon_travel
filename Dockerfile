FROM golang:1.16.3-alpine3.13

RUN apk add --no-cache git gcc libc-dev

WORKDIR /usr/local/go/src/awesomeProject
RUN go get -u github.com/tidwall/gjson
COPY . .
RUN go mod tidy && go mod vendor

