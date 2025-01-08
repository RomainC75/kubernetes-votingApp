
# -f to change the context
build-voting-server: 
	docker build --progress=plain -t voting-server -f ./voting-server/Dockerfile .