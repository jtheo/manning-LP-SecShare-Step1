# Secrets Sharing
Manning liveProject "(Build a Secrets Sharing Web Application)[https://www.manning.com/liveproject/build-a-secrets-sharing-web-application]" 

    Milestone 1: you will implement a web application that allow creating a new secret and viewing the secret. Once viewed, a secret should not be viewable again. You will use a file to store the secrets. You will verify the behavior by manually making HTTP requests using a program such as curl.

**Test**
```
$: curl -X POST http://localhost:8080 -d '{"plain_text":"My super secret123"}'
{"id": "c616584ac64a93aafe1c16b6620f5bcd"}

$: curl http://localhost:8080/c616584ac64a93aafe1c16b6620f5bcd
{"data": "My super secret123"}

$: curl http://localhost:8080/c616584ac64a93aafe1c16b6620f5bcd
{"data": ""}

$: curl http://localhost:8080/ -i
HTTP/1.1 400 Bad Request
Date: Wed, 08 Sep 2021 12:03:58 GMT
Content-Length: 18
Content-Type: text/plain; charset=utf-8

Bad, Bad Request!


$: curl http://localhost:8080/ -i -d 'Something'
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Date: Wed, 08 Sep 2021 12:04:08 GMT
Content-Length: 12

Unknown Key
```


I know the code it's a mess, please be gentle :)


**PS**
Do not use this code for anything unless it's a local test.