FROM golang:1.23.3-alpine

RUN go install github.com/air-verse/air@latest

RUN apk add --no-cache postgresql-client make
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /go/src/voting-server

COPY custom-workspaces/go.voting-server.work /go/src/go.work
COPY ./shared /go/src/shared
COPY ./voting-server .
RUN ls
RUN go mod download

# RUN pwd
# RUN ls /
# COPY ./docker/local/server/go.work /go/src/
# RUN ls

# CMD ["air"]
CMD echo "=======STARTUP=======" && ls && sleep 10000 && printenv && migrate --path server/db/migration --database "postgresql://name:pass@localhost:5432/securitest?sslmode=disable" --verbose up && air