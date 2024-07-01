# SSO

SSO service

## Stack:

gRPC - google.golang.org/grpc

Generated code - github.com/paniccaaa/protos (repo which contains proto files and gen. code for gRPC and protobuf)

PostgreSQL driver - github.com/jackc/pgx/v5 

Lib for jwt - github.com/golang-jwt/jwt/v5

Enviroment config reader - github.com/ilyakaznacheev/cleanenv

Logging - log/slog (default package)

Hashing passwords - golang.org/x/crypto/bcrypt

Migrate- github.com/golang-migrate/migrate/v4

## Database migration (check **cmd/migrator**)

You can now perform database migrations without installing the [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) utility. Use the following commands

```bash
go run ./cmd/migrator --db-uri="postgres://<username>:<password>@<host>:<port>/<database_name>" --migrations-path=./migrations --direction=up

# and

go run ./cmd/migrator --db-uri="postgres://<username>:<password>@<host>:<port>/<database_name>" --migrations-path=./migrations --direction=down
```