FROM golang:1.23.3-alpine

RUN go install github.com/air-verse/air@latest

RUN apk add --no-cache postgresql-client make
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

WORKDIR /go/src/job

COPY custom-workspaces/go.job.work /go/src/go.work
COPY ./shared /go/src/shared
COPY ./job .

RUN ls
RUN go mod download

# RUN pwd
# RUN ls /
# COPY ./docker/local/server/go.work /go/src/
# RUN ls

# CMD ["air"]
CMD ls && go run main.go