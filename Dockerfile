FROM golang:1.19

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client
RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o shortener ./cmd/shortener/main.go

CMD ["./shortener"]
