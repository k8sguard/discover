build-image:
	docker build -f ./hack/images/Dockerfile -t local/discover .