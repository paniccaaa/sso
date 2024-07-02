# SSO

SSO service 

Protobuf contract - https://github.com/paniccaaa/protos

## Stack:

- gRPC

- PostgreSQL

- JWT 

- cleanevn (env. config reader)

- slog (logger)

## Installation 

1. **Clone the repo**

```bash
git clone https://github.com/paniccaaa/sso.git
cd sso
```

2. **local.yaml and .env files**

Make sure to fill out the local.yaml and .env files with the necessary configuration settings.

3. **Run the app locally**

```bash
make run
```

## Database migration (check **cmd/migrator**)

You can now perform database migrations without installing the [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) utility. Use the following commands

```bash
go run ./cmd/migrator --db-uri="postgres://<username>:<password>@<host>:<port>/<database_name>" --migrations-path=./migrations --direction=up

# and

go run ./cmd/migrator --db-uri="postgres://<username>:<password>@<host>:<port>/<database_name>" --migrations-path=./migrations --direction=down
```