FROM golang:1.16.3-alpine
RUN apk add --update &&  apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app

RUN go mod init quickstart
# RUN go get google.golang.org/api/gmail/v1
# RUN go get golang.org/x/oauth2/google
