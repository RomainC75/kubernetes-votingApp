FROM golang:1.23.3-alpine

RUN go install github.com/air-verse/air@latest

WORKDIR /go/src/front-server

COPY ./front-server .

RUN ls
RUN go mod download

CMD ls && air