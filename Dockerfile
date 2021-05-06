FROM golang:1.16.3-alpine3.13
RUN apk add --no-cache git gcc libc-dev
WORKDIR ./app
COPY . .
RUN go mod init github.com/DaniilMats/ozon_travel && go mod tidy

