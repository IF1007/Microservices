FROM golang:1.11

ENV GOOS=linux

WORKDIR $GOPATH/src/publisher

COPY . .

ENV GO111MODULE=on

RUN go mod vendor
RUN go build -mod=vendor
RUN chmod +x ./dcm-publisher

CMD ["./dcm-publisher"]
