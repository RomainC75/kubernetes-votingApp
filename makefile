INFRA_FOLDER=infra

build-voting-server:
	docker build --progress=plain -t voting-server -f ./voting-server/Dockerfile .

build-job: 
	docker build --progress=plain -t job -f ./job/Dockerfile .

build-front-server: 
	docker build --progress=plain -t front-server -f ./front-server/Dockerfile .

# ========= K8S =========

reload-voting-server:
	kubectl delete -f $(INFRA_FOLDER)/voting-server.yaml; \
	kubectl apply -f $(INFRA_FOLDER)/voting-server.yaml

reload-front-server:
	kubectl delete -f $(INFRA_FOLDER)/front-server.yaml; \
	kubectl apply -f $(INFRA_FOLDER)/front-server.yaml

reload-postgres:
	kubectl delete -f $(INFRA_FOLDER)/postgres.yaml; \
	rm -rf data/db; \
	kubectl apply -f $(INFRA_FOLDER)/postgres.yaml

reload-redis:
	kubectl delete -f $(INFRA_FOLDER)/redis.yaml; \
	kubectl apply -f $(INFRA_FOLDER)/redis.yaml

reload-job:
	kubectl delete -f $(INFRA_FOLDER)/job.yaml; \
	kubectl apply -f $(INFRA_FOLDER)/job.yaml

reload-configs:
	kubectl delete -f $(INFRA_FOLDER)/configs.yaml; \
	kubectl create -f $(INFRA_FOLDER)/configs.yaml;

reload-secrets:
	kubectl delete -f $(INFRA_FOLDER)/secrets.yaml; \
	kubectl create -f $(INFRA_FOLDER)/secrets.yaml;

build-all: build-voting-server build-job build-front-server
run-data: reload-configs reload-secrets
run-dbs: reload-redis reload-postgres
run-servers: reload-voting-server reload-front-server reload-job

remove-everything:
	kubectl delete -f $(INFRA_FOLDER)/job.yaml; \
	kubectl delete -f $(INFRA_FOLDER)/front-server.yaml; \
	kubectl delete -f $(INFRA_FOLDER)/voting-server.yaml; \
	kubectl delete -f $(INFRA_FOLDER)/postgres.yaml; \
	kubectl delete -f $(INFRA_FOLDER)/redis.yaml; \
	kubectl delete -f $(INFRA_FOLDER)/configs.yaml; \
	kubectl delete -f $(INFRA_FOLDER)/secrets.yaml;

sqlc:
	cd voting-server && sqlc generate && ./comment-cleaner.sh 