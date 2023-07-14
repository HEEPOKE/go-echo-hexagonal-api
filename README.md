# go-echo-hexagonal-api

## Run DB With Docker

```bash
docker create network Heepoke
```

```bash
docker compose up -d
```

## config .env

```bash
cp .env.example .env
```

- then update environments in .env

## run

```bash
go mod tidy
```

```bash
go run cmd/main.go
```

## Generate Swagger

```bash
swag init -g cmd/main.go --output=internal/app/docs
```

## Local Swagger Doc Api

```bash
http://localhost:6476/apis/docs/index.html
```
