# Task Manager

```sh
go mod init task_manager
go get -u github.com/gin-gonic/gin
go get github.com/redis/go-redis/v9
```

## Redis

```sh
docker compose up
docker exec -it task-manager-redis-1 bash
ls -al /usr/local/bin/  # Находим redis-cli здесь
redis-cli
ZRANGE tasks 0 -1 WITHSCORES
> (empty array)
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
