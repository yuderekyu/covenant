FROM golang

ADD ./covenant /go/bin/covenant
ADD ./config.json /go/bin/config.json

ENTRYPOINT /go/bin/covenant

EXPOSE 8082
