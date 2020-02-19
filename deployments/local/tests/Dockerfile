FROM golang:1.13

RUN go get -d -u github.com/Shodocan/InstanceInventoryApi ; exit 0

ENV DB_DATABASE Instance
ENV DB_PORT 27017
ENV DB_HOST db
ENV LOG_LEVEL Debug

WORKDIR /go/src/github.com/Shodocan/InstanceInventoryApi

RUN go get -t ./...

ENTRYPOINT [ "go","test", "./...", "-timeout=2m", "-parallel=4",  "-covermode=atomic", "-coverprofile", "coverage.out" ]