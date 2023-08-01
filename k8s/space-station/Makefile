
.PHONY:
create-cluster:
	k3d cluster create space-station-cluster --api-port 6551 --servers 1 --agents 1

.PHONY:
delete-cluster:
	k3d cluster delete space-station-cluster

.PHONY:
build-image:
	docker build ./app -f ./k8s/Dockerfile -t space-station:local
	
.PHONY:
import-image:
	k3d image import -c space-station-cluster space-station:local

.PHONY:
deploy:
	kubectl apply -f ./k8s/space-station-deploy.yaml

.PHONY:
all: create-cluster build-image import-image deploy