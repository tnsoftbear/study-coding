### Ok

GET http://localhost:8080/sum/123/32
Origin: http://any-origin.io
Content-type: test
HTTP 200
content-type: application/json
content-length: 16
"123 + 32 = 155"

# 404: Route not found

GET http://localhost:8080/unknown
HTTP 404
[Asserts]
body == "Route not found"

GET http://localhost:8080/sum/123/a32
HTTP 422
[Asserts]
body == "Problem with the second value (\"a32\") - invalid digit found in string"
