FROM golang:1.11

ENV GOOS=linux

WORKDIR $GOPATH/src/github.com/esvm/microservices

COPY . .

ENV GO111MODULE=on

RUN go mod download
RUN go build -o consumer src/consumer/consumer.go
RUN chmod +x ./consumer

CMD ["./consumer"]

