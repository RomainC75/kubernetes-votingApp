
# -f to change the context
build-voting-server: 
	docker build --progress=plain -t voting-server -f ./voting-server/Dockerfile .

reload-voting-server:
	kubectl delete -f infra/voting-server.yaml; \
	kubectl apply -f infra/voting-server.yaml

reload-postgres:
	kubectl delete -f infra/db.yaml; \
	rm -rf data/db; \
	kubectl apply -f infra/db.yaml