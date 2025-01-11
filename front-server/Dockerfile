FROM golang:1.23.3-alpine

RUN go install github.com/air-verse/air@latest

WORKDIR /go/src/front-server


COPY custom-workspaces/go.front-server.work /go/src/go.work
COPY ./shared /go/src/shared
COPY ./front-server .

RUN ls
RUN go mod download

CMD ls && air