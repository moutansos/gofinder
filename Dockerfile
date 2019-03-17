FROM golang:1.12

ENV GO111MODULE=on

WORKDIR /go/src/gofinder
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

VOLUME /srv
WORKDIR /srv

CMD ["gofinder"]