FROM golang:buster as builder

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /go/src/app

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

CMD air