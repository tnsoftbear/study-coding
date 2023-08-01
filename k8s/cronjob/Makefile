
.PHONY:
create-cluster:
	k3d cluster create cj-cluster --api-port 6550 --servers 1 --agents 1

.PHONY:
delete-cluster:
	k3d cluster delete cj-cluster

.PHONY:
deploy:
	kubectl apply -f ./cronjob.yaml
