### Ok: Ping-pong
GET http://localhost:8080/ping
Origin: http://any-origin.io
Content-type: any-content-type
HTTP 200
content-type: application/json; charset=utf-8
content-length: 18
{"message":"pong"}

### Fail: 404: Route not found
GET http://localhost:8080/unknown
HTTP 404
[Asserts]
body == "404 page not found"

### Ok: Add task
POST http://localhost:8080/task
Content-type: application/json
{"id":"task-id-1","name":"task-name-1","description":"task-description-1","created_at":1000001}
HTTP 201
content-type: application/json; charset=utf-8
content-length: 157
{"created":true,"message":"Task created successfully","task":{"id":"task-id-1","name":"task-name-1","description":"task-description-1","created_at":1000001}}

POST http://localhost:8080/task
Content-type: application/json
{"id":"task-id-2","name":"task-name-2","description":"task-description-2","created_at":1000002}
HTTP 201
content-type: application/json; charset=utf-8
content-length: 157
{"created":true,"message":"Task created successfully","task":{"id":"task-id-2","name":"task-name-2","description":"task-description-2","created_at":1000002}}

### Ok: Get task by id
GET http://localhost:8080/task/task-id-1
Content-type: application/json
HTTP 200
content-type: application/json; charset=utf-8
content-length: 104
{"task":{"id":"task-id-1","name":"task-name-1","description":"task-description-1","created_at":1000001}}

### Ok: Get task list
GET http://localhost:8080/task
Content-type: application/json
HTTP 200
content-type: application/json; charset=utf-8
content-length: 437
{"tasks":[{"id":"task-id-1","name":"task-name-1","description":"task-description-1","created_at":1000001},{"id":"task-id-2","name":"task-name-2","description":"task-description-2","created_at":1000002},{"id":"sample-task-id-1","name":"Sample task 1 name","description":"Sample task 1 description","created_at":2000001},{"id":"sample-task-id-2","name":"Sample task 2 name","description":"Sample task 2 description","created_at":2000002}]}

### Fail: Delete task by id
DELETE http://localhost:8080/task/absent-task-id
Content-type: application/json
HTTP 404
content-type: application/json; charset=utf-8
content-length: 88
{"deleted":false,"id":"absent-task-id","message":"Task not found by id: absent-task-id"}

### Ok: Delete task by id
DELETE http://localhost:8080/task/task-id-1
Content-type: application/json
HTTP 200
content-type: application/json; charset=utf-8
content-length: 71
{"deleted":true,"id":"task-id-1","message":"Task successfully deleted"}

### Fail: Get task by id
GET http://localhost:8080/task/task-id-1
Content-type: application/json
HTTP 404
content-type: application/json; charset=utf-8
content-length: 62
{"id":"task-id-1","message":"Task not found by id: task-id-1"}
