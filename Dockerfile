FROM golang:1.16-alpine

RUN apk add --update curl bash git

WORKDIR /go/src

COPY . .

RUN go get github.com/codegangsta/gin
RUN go mod tidy

EXPOSE 8080

ENTRYPOINT [ "top" ]