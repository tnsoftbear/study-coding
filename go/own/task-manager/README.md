# Task Manager

```sh
docker compose up
docker exec -it task-manager-redis-1 bash
docker run --rm -it --entrypoint bash redis -c "redis-cli -h host.docker.internal -p 6379"
# Соберём апку в докере и запустим с подключение к редис контейнеру
docker build -f .\\infra\\docker\\app\\Dockerfile -t task-manager:0.1 .
docker run --rm -p 8080:8080 --env REDIS_HOST=host.docker.internal:6379 task-manager:0.1
# Посмотрим сети докера, а так же сеть контейнера redis
docker network ls
docker inspect task-manager-redis-1
# Подключим контейнер к сети task-manager_default и укажем имя службы "redis:6379" для хоста редис, которое будет преобразовано в IP внутри сети:
docker run --rm -p 8080:8080 --env REDIS_HOST=redis:6379 --network=task-manager_default task-manager:0.1
```

Далее я перенёс эту процедуру сборки в `docker-compose.yaml` добавив сервис `app`

```yaml
  app:
    build: 
      context: .
      dockerfile: ./infra/docker/app/Dockerfile
    ports:
      - 8080:8080
    environment:
      - REDIS_HOST=redis:6379
```

## Testing

```sh
# install hurl at first (cargo install hurl)
hurl ./test/web_server.hurl
# or manually
curl -X POST http://localhost:8080/task -H "Content-type: application/json" -d "{\"id\":\"task-id-1\",\"name\":\"task-name-1\",\"description\":\"task-description-1\",\"created_at\":100000}"
```

## Links

* [gin quickstart](https://gin-gonic.com/docs/quickstart/)
* [template](https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go)
* [gin docs](https://pkg.go.dev/github.com/gin-gonic/gin)
* [go-redis](https://github.com/redis/go-redis)
