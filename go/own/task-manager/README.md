# Task Manager

## Install

```sh
docker compose up -d
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
