### Ok: Ping-pong
GET {{url_base}}/ping
Origin: http://any-origin.io
Content-type: any-content-type
HTTP 200
content-type: application/json
content-length: 18
{"message":"pong"}

### Fail: 404: Route not found
GET {{url_base}}/unknown
HTTP 404
[Asserts]
body == "404 page not found"

### Ok: Successful insert 1
POST {{url_base}}/parcel-locker
Content-type: application/json
{"id":"parcel-locker-id-1","name":"The Olympic Sports Centre","longitude":24.124542,"latitude":56.967747}
HTTP 201
content-type: application/json
content-length: 105
{"id":"parcel-locker-id-1","name":"The Olympic Sports Centre","longitude":24.124542,"latitude":56.967747}

### Ok: Successful update 1
POST {{url_base}}/parcel-locker
Content-type: application/json
{"id":"parcel-locker-id-1","name":"The \"Olympic\" Sports Centre","longitude":24.111111,"latitude":56.999999}
HTTP 200
content-type: application/json
content-length: 109
{"id":"parcel-locker-id-1","name":"The \"Olympic\" Sports Centre","longitude":24.111111,"latitude":56.999999}

### Ok: Successful insert 2
POST {{url_base}}/parcel-locker
Content-type: application/json
{"id":"parcel-locker-id-2","name":"The Olympic Sports Centre 2","longitude":24.324542,"latitude":57.012123}
HTTP 201
content-type: application/json
content-length: 107
{"id":"parcel-locker-id-2","name":"The Olympic Sports Centre 2","longitude":24.324542,"latitude":57.012123}

### Ok: Successful insert 3
POST {{url_base}}/parcel-locker
Content-type: application/json
{"id":"parcel-locker-id-3","name":"The Olympic Sports Centre 3","longitude":25.00,"latitude":58.00}
HTTP 201
content-type: application/json
content-length: 97
{"id":"parcel-locker-id-3","name":"The Olympic Sports Centre 3","longitude":25.0,"latitude":58.0}

### Ok: Read one
GET {{url_base}}/parcel-locker/parcel-locker-id-1
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 109
{"id":"parcel-locker-id-1","name":"The \"Olympic\" Sports Centre","longitude":24.111111,"latitude":56.999999}

### Ok: Distance search
GET {{url_base}}/parcel-locker/distance-search?latitude=57.01&longitude=24.20&radius=10
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 159
[{"name":"parcel-locker-id-1","latitude":0.0,"longitude":0.0,"distance":5.4978},{"name":"parcel-locker-id-2","latitude":0.0,"longitude":0.0,"distance":7.5461}]

### Ok: Distance search nothing found
GET {{url_base}}/parcel-locker/distance-search?latitude=57.01&longitude=24.20&radius=1
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 2
[]

### Ok: Delete existing entry
DELETE {{url_base}}/parcel-locker/parcel-locker-id-1
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 176
{"deleted":true,"message":"Parcel locker deleted","parcel_locker":{"id":"parcel-locker-id-1","name":"The \"Olympic\" Sports Centre","longitude":24.111111,"latitude":56.999999}}

### Fail: Delete absent entry 1
DELETE {{url_base}}/parcel-locker/parcel-locker-id-1
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 100
{"deleted":false,"message":"Parcel locker not found by id: parcel-locker-id-1","parcel_locker":null}

### Ok: Distance search again
GET {{url_base}}/parcel-locker/distance-search?latitude=57.01&longitude=24.20&radius=10
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 80
[{"name":"parcel-locker-id-2","latitude":0.0,"longitude":0.0,"distance":7.5461}]

### Ok: Delete existing entry 2
DELETE {{url_base}}/parcel-locker/parcel-locker-id-2
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 174
{"deleted":true,"message":"Parcel locker deleted","parcel_locker":{"id":"parcel-locker-id-2","name":"The Olympic Sports Centre 2","longitude":24.324542,"latitude":57.012123}}

### Ok: Delete existing entry 3
DELETE {{url_base}}/parcel-locker/parcel-locker-id-3
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 164
{"deleted":true,"message":"Parcel locker deleted","parcel_locker":{"id":"parcel-locker-id-3","name":"The Olympic Sports Centre 3","longitude":25.0,"latitude":58.0}}
