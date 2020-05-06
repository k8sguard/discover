build-image:
	docker build -f ./hack/images/Dockerfile -t local/discover .
depoloy-minikube:
	kubectl apply -f deploy/deploy-minikube.yaml