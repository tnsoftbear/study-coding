cf = -f deploy/docker/compose.yaml

help: ## Print this help
	@grep -E '^[a-zA-Z_-]+:.*## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Run in docker 

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

appup:
	@go run ./cmd/
doload:
	@for i in {1..100}; do curl 0x0:5000/do; done
metric1:
	@curl -s 0x0:8000 | grep "tp_req_total"
metric2:
	@curl -s 0x0:8000 | grep "tp_in_progress_req_count"
metric3:
	@curl -s 0x0:8000 | grep "tp_req_latency"

.PHONY: \
		appup \
		build \
		doload \
		down \
		help \
		metric1 \
		metric2 \
		metric3 \
		rebuild \
		restart \
		up \
