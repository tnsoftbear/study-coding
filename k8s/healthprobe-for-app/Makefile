
.PHONY:
create-cluster:
	k3d cluster create health-check-cluster --api-port 6550 --servers 1 --agents 1

.PHONY:
delete-cluster:
	k3d cluster delete health-check-cluster

.PHONY:
build-image:
	docker build ./app -f ./docker/Dockerfile -t health-app:local
	
.PHONY:
import-image:
	k3d image import -c health-check-cluster health-app:local

.PHONY:
deploy:
	kubectl apply -f ./k8s/deployment.yaml

.PHONY:
build-and-deploy:
	make build-image
	make import-image
	make deploy