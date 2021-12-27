FROM golang:1.15.2-buster

COPY . /go/src/gitlab.com/andersph/docker-master-class/api
WORKDIR /go/src/gitlab.com/andersph/docker-master-class/api

ENTRYPOINT [ "bash", "-c", "go run main.go" ]