FROM golang

ADD . /go/src/github.com/user_name/project_name

WORKDIR /go/src/github.com/user_name/project_name

RUN go install github.com/user_name/deploy_docker

ENTRYPOINT /go/bin/deploy_docker

EXPOSE 8082