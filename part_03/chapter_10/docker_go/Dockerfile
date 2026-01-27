FROM golang

ADD . /go/src/github.com/user_name/docker_go

WORKDIR /go/src/github.com/user_name/docker_go

RUN go get github.com/lib/pq
RUN go install github.com/user_name/docker_go

ENTRYPOINT /go/bin/docker_go

EXPOSE 8082