### Ok: Ping-pong
GET {{url_base}}/ping
Origin: http://any-origin.io
Content-type: any-content-type
HTTP 200
content-type: application/json; charset=utf-8
content-length: 18
{"message":"pong"}

### Fail: 404: Route not found
GET {{url_base}}/unknown
HTTP 404
[Asserts]
body == "404 page not found"

### Ok: Add shipping
POST {{url_base}}/customer/shipping
Content-type: application/json
{"id":"shipping-id-1","name":"shipping-name-1","created_at":1000001,"address":{"id":"shipping-address-id-1","city":"Riga","street":"Brivibas 123","longitude":24.11,"latitude":56.01}}
HTTP 201
content-type: application/json; charset=utf-8
content-length: 269
{"created":true,"message":"Customer shipping address created successfully","shipping":{"id":"shipping-id-1","name":"shipping-name-1","address":{"id":"shipping-address-id-1","city":"Riga","street":"Brivibas 123","longitude":24.11,"latitude":56.01},"created_at":1000001}}

POST {{url_base}}/customer/shipping
Content-type: application/json
{"id":"shipping-id-2","name":"shipping-name-2","created_at":1000002}
HTTP 201
content-type: application/json; charset=utf-8
content-length: 170
{"created":true,"message":"Customer shipping address created successfully","shipping":{"id":"shipping-id-2","name":"shipping-name-2","address":null,"created_at":1000002}}

### Ok: Get shipping by id
GET {{url_base}}/customer/shipping/shipping-id-1
Content-type: application/json
HTTP 200
content-type: application/json; charset=utf-8
content-length: 196
{"shipping":{"id":"shipping-id-1","name":"shipping-name-1","address":{"id":"shipping-address-id-1","city":"Riga","street":"Brivibas 123","longitude":24.11,"latitude":56.01},"created_at":1000001}}

### Ok: Get shipping list
GET {{url_base}}/customer/shippings
Content-type: application/json
HTTP 200
content-type: application/json; charset=utf-8
content-length: 282
{"shippings":[{"id":"shipping-id-1","name":"shipping-name-1","address":{"id":"shipping-address-id-1","city":"Riga","street":"Brivibas 123","longitude":24.11,"latitude":56.01},"created_at":1000001},{"id":"shipping-id-2","name":"shipping-name-2","address":null,"created_at":1000002}]}

### Fail: Delete shipping by id
DELETE {{url_base}}/customer/shipping/absent-shipping-id
Content-type: application/json
HTTP 404
content-type: application/json; charset=utf-8
content-length: 117
{"deleted":false,"id":"absent-shipping-id","message":"Customer shipping address not found by id: absent-shipping-id"}

### Ok: Delete shipping by id
DELETE {{url_base}}/customer/shipping/shipping-id-1
Content-type: application/json
HTTP 200
content-type: application/json; charset=utf-8
content-length:96
{"deleted":true,"id":"shipping-id-1","message":"Customer shipping address successfully deleted"}

### Fail: Get shipping by id
GET {{url_base}}/customer/shipping/shipping-id-1
Content-type: application/json
HTTP 404
content-type: application/json; charset=utf-8
content-length: 91
{"id":"shipping-id-1","message":"Customer shipping address not found by id: shipping-id-1"}
