# ECDH Master
It's a repository for [ECDH Master](https://github.com/artemioo/ecdhsnap) backend.

## Tools&Libs
- pgx/stdlib v5.5.2
- sqlx v1.3.5
- squirrel v1.5.4
- migrate v4.17.1
- viper v1.18.2
- gotenv v1.6.0
- go-chi/cors v1.2.1
- gorilla/sessions v1.2.2

## Getting Started
You need to run migrations. For example:

```shell
sudo migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' -verbose up`
```
NOTE: paste here your variables

And run at main directory:
```shell
go run cmd/main.go
```