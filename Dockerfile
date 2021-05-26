FROM golang:latest
ADD . /go/src/github.com/suderio/gosh
WORKDIR /go/src/github.com/suderio/gosh
RUN go install -v
ENTRYPOINT ["gosh"]