## build voting-server

docker build -t voting-server ./voting-server/

## DB 

DB_USERNAME=my-name
DB_PASSWORD=my-pass
DB_NAME=mydb
!! delete the data folder to reset the db !!

migrate --path db/migration --database "postgresql://my-name:my-pass@postgres-cluster-srv.default.svc.cluster.local:5432/vote-db?sslmode=disable" --verbose up