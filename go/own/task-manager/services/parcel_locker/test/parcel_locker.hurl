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
{"error":"404 page not found"}

### Ok: Clear DB
DELETE {{url_base}}/all-parcel-lockers
HTTP 200

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

### Fail: Missing field insert
POST {{url_base}}/parcel-locker
Content-type: application/json
{"id":"parcel-locker-id-2","name":"The Olympic Sports Centre","longitude":24.124542}
HTTP 422
content-type: application/json
content-length: 88
{"error":"Request body deserialize error: missing field `latitude` at line 1 column 84"}

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

### Ok: Read all
GET {{url_base}}/parcel-lockers
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 317
[{"id":"parcel-locker-id-1","name":"The \"Olympic\" Sports Centre","longitude":24.111111,"latitude":56.999999},{"id":"parcel-locker-id-2","name":"The Olympic Sports Centre 2","longitude":24.324542,"latitude":57.012123},{"id":"parcel-locker-id-3","name":"The Olympic Sports Centre 3","longitude":25.0,"latitude":58.0}]

### Ok: Read first page
GET {{url_base}}/parcel-lockers?page=1&per_page=2
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 219
[{"id":"parcel-locker-id-1","name":"The \"Olympic\" Sports Centre","longitude":24.111111,"latitude":56.999999},{"id":"parcel-locker-id-2","name":"The Olympic Sports Centre 2","longitude":24.324542,"latitude":57.012123}]

### Ok: Read second page
GET {{url_base}}/parcel-lockers?page=2&per_page=2
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 99
[{"id":"parcel-locker-id-3","name":"The Olympic Sports Centre 3","longitude":25.0,"latitude":58.0}]

### Ok: Read third page
GET {{url_base}}/parcel-lockers?page=3&per_page=2
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 2
[]

### Ok: Distance search
GET {{url_base}}/parcel-locker-distance-search?latitude=57.01&longitude=24.20&distance=10
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 216
[{"name":"parcel-locker-id-1","latitude":56.99999958780393,"longitude":24.11110907793045,"distance":5.4978},{"name":"parcel-locker-id-2","latitude":57.01212315910904,"longitude":24.324543178081512,"distance":7.5461}]

### Ok: Distance search nothing found
GET {{url_base}}/parcel-locker-distance-search?latitude=57.01&longitude=24.20&distance=1
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 2
[]

### Fail: Distance search invalid "longitude" parameter
GET {{url_base}}/parcel-locker-distance-search?latitude=57.01&longitude=24x20&distance=1
Content-type: application/json
HTTP 400
content-type: application/json
content-length: 57
{"error":"\"longitude\" parameter must be numeric value"}

### Fail: Distance search invalid "latitude" parameter
GET {{url_base}}/parcel-locker-distance-search?latitude=57.01x&longitude=24.20&distance=1
Content-type: application/json
HTTP 400
content-type: application/json
content-length: 56
{"error":"\"latitude\" parameter must be numeric value"}

### Fail: Distance search invalid "distance" parameter
GET {{url_base}}/parcel-locker-distance-search?latitude=57.01&longitude=24.20&distance=x1
Content-type: application/json
HTTP 400
content-type: application/json
content-length: 54
{"error":"\"distance\" parameter must be numeric value"}

### Fail: Distance search missing "longitude" parameter
GET {{url_base}}/parcel-locker-distance-search?latitude=57.01&distance=1
Content-type: application/json
HTTP 400
content-type: application/json
content-length: 44
{"error":"\"longitude\" parameter required"}

### Fail: Distance search missing "latitude" parameter
GET {{url_base}}/parcel-locker-distance-search?longitude=24.20&distance=1
Content-type: application/json
HTTP 400
content-type: application/json
content-length: 43
{"error":"\"latitude\" parameter required"}

### Fail: Distance search missing "distance" parameter
GET {{url_base}}/parcel-locker-distance-search?latitude=57.01&longitude=24.20
Content-type: application/json
HTTP 400
content-type: application/json
content-length: 41
{"error":"\"distance\" parameter required"}

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
GET {{url_base}}/parcel-locker-distance-search?latitude=57.01&longitude=24.20&distance=10
Content-type: application/json
HTTP 200
content-type: application/json
content-length: 109
[{"name":"parcel-locker-id-2","latitude":57.01212315910904,"longitude":24.324543178081512,"distance":7.5461}]

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
