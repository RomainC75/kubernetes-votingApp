FROM golang:1.23.3-alpine

RUN go install github.com/air-verse/air@latest

RUN apk add --no-cache postgresql-client make
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /go/src/server

COPY . .
# COPY go.mod go.sum ./
RUN go mod download


RUN pwd

RUN ls /

# COPY ./docker/local/server/go.work /go/src/

RUN ls

# CMD ["air"]
CMD echo "=======STARTUP=======" && migrate -path db/migrations -database "$DATABASE_URL" up && air