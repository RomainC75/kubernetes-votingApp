
# -f to change the context
build-voting-server: 
	docker build --progress=plain -t voting-server -f ./voting-server/Dockerfile .

build-job: 
	docker build --progress=plain -t job -f ./job/Dockerfile .

build-front-server: 
	docker build --progress=plain -t front-server -f ./front-server/Dockerfile .

reload-voting-server:
	kubectl delete -f infra/voting-server.yaml; \
	kubectl apply -f infra/voting-server.yaml

reload-front-server:
	kubectl delete -f infra/front-server.yaml; \
	kubectl apply -f infra/front-server.yaml

reload-postgres:
	kubectl delete -f infra/postgres.yaml; \
	rm -rf data/db; \
	kubectl apply -f infra/postgres.yaml

reload-job:
	kubectl delete -f infra/job.yaml; \
	kubectl apply -f infra/job.yaml

sqlc:
	cd voting-server && sqlc generate && ./comment-cleaner.sh 