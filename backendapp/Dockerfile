FROM golang:1.10.0-stretch

RUN apt-get update && \
    apt-get install -y postgresql-client
RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/gaku3601/gApiG/backendapp
COPY . .

RUN dep ensure 
RUN go install github.com/gaku3601/gApiG/backendapp

CMD bash ./start.sh
