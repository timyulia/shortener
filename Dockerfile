#FROM golang:1.19
#
##ENV MODE=inMemory
#ENV MODE=db
#
#RUN go version
#ENV GOPATH=/
#
#COPY ./ ./
#
#RUN apt-get update
#RUN apt-get -y install postgresql-client
#RUN chmod +x wait-for-postgres.sh
#
#
#RUN go mod download
#RUN go build -o shortener ./cmd/shortener/main.go
#
#RUN echo TESTEST
#
#CMD ["./shortener"]

FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go get -u github.com/pressly/goose/cmd/goose
RUN goose -version

COPY . .

RUN go build -o main ./cmd/shortener/main.go

CMD ["./main"]