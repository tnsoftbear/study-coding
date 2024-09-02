cf = -f deploy/docker/compose.yaml

build: ## Build docker containers
	docker compose $(cf) build
up: ## Start docker containers
	docker compose $(cf) up -d --remove-orphans
down: ## Stop docker containers
	docker compose $(cf) down
rebuild: ## Rebuild and start docker containers
	@make down
	@make build
	@make up
restart: ## Restart
	docker compose $(cf) restart

# Local development

gen: ## Generate code for reform logic
	go generate ./...

dockerinit:
	docker network create frr
dbup:	# https://hub.docker.com/_/mysql
	docker run --name frr-news-storage -p3307:3306 --network frr \
	-e MYSQL_ROOT_PASSWORD=pw -e MYSQL_DATABASE=frr -e MYSQL_USER=admin -e MYSQL_PASSWORD=123 \
	-v E:/coding/my/go/own/fiber-reform-rest/deploy/docker/storage/initdb:/docker-entrypoint-initdb.d:ro \
	-d mysql:latest
dbdown:
	docker stop frr-news-storage
	docker rm frr-news-storage
dbrestart:
	@make dbdown
	@make dbup
dbclient:
	docker run -it --rm --network frr mysql mysql -hfrr-news-storage -uadmin -p --database frr
localdbclient:
	mysql.exe --port 3307 -uadmin -p --database frr

.PHONY: \
		build \
		down \
		rebuild \
		up \