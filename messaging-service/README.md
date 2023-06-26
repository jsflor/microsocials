# GOLANG MICROSERVICE

It uses Gin and Gorm

## Install dependencies

```bash
go mod tidy
```

## Test app

```bash
go test -v ./..
```

## Run service

```bash
go run .
```

## Build service

```bash
go build .
```

## Build docker image from given Dockerfile

```bash
docker build -t jsflor/messaging-service:1.0 .
```

## Run docker image exposing container port

```bash
docker run -p 8080:8080 jsflor/messaging-service:1.0
```
