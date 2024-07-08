hurl:
	hurl --variables-file=.\test\variables .\test\parcel_locker.hurl
build:
	docker compose -f deploy/docker/docker-compose.yaml build
up:
	docker compose -f deploy/docker/docker-compose.yaml up -d
down:
	docker compose -f deploy/docker/docker-compose.yaml down