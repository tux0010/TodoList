FROM golang:latest

COPY ./Todo /go/bin/Todo
ENTRYPOINT /go/bin/Todo

EXPOSE 8080
