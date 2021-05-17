FROM golang:latest

COPY . /my_go/src/rest-api

WORKDIR /my_go/src/rest-api/cmd/rest-api

RUN go build -o rest-api main.go

EXPOSE 8080

CMD ["./rest-api"]