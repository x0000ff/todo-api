# Simple CRUD API for TODO app

![[golang](https://golang.org)](https://img.shields.io/badge/golang-1.8-blue.svg?style=flat-square)

## How to use

```terminal
$ go run main.go
```

### Index 

```terminal
curl -X GET \
  http://localhost:8080/api/v1/todos/ \
  -H 'cache-control: no-cache'
```

### Create 

```terminal
curl -X POST \
  http://localhost:8080/api/v1/todos/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F title=Second \
  -F completed=false
```

### Show 

```terminal
curl -X GET \
  http://localhost:8080/api/v1/todos/2
  -H 'cache-control: no-cache'
```

### Update 

```terminal
curl -X PUT \
  http://localhost:8080/api/v1/todos/2 \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F title=Second \
  -F completed=true
```

### Delete 

```terminal
curl -X DELETE \
  http://localhost:8080/api/v1/todos/7 \
  -H 'cache-control: no-cache'
```

## Used materials

- 🇬🇧 [Build RESTful API service in golang using gin-gonic framework](https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3)
- 🇬🇧 [Developing a simple CRUD API with Go, Gin and Gorm](https://medium.com/@cgrant/developing-a-simple-crud-api-with-go-gin-and-gorm-df87d98e6ed1)
- 🇷🇺 [Создания API RESTful сервера с Golang и MySQL](https://medium.com/go-to-golang/%D1%81%D0%BE%D0%B7%D0%B4%D0%B0%D0%BD%D0%B8%D1%8F-api-restful-%D1%81%D0%B5%D1%80%D0%B2%D0%B5%D1%80%D0%B0-%D1%81-golang-%D0%B8-mysql-3085133bab91)
- 🇷🇺 [Язык Go для начинающих](https://habr.com/post/219459/)

## Credits

- [GORM](http://gorm.io)
- [Gin](https://github.com/gin-gonic/gin)
