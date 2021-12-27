# README

## Postgres Preparation

* migration in postgres container:
```psql -U admin -d my_company < /src/migration/0000_initial_migration.sql```

## common go commands

### run and test go programm

```bash
go run main.go
```

### build binary

```bash
go build
```

### run test

```bash
go test -v
```

### create module

```bash
go mod init
```

### clean required modules

```bash
go mod tidy
```

### load vendor packages

With the execution of this command vendor packages do not have to be downloaded
with every build and therefore have a static version.

```bash
go mod vendor
```

### curl test example

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Max Mustermann","address":"Musterstreet", "email":"max@mustermann.org", "birth":"11.11.2000", "department":"Test", "job_title":"unknown"}' \
  http://localhost:8080/employee
```  

## Environment variables

Following environment variables can be set for the docker image:

* ENV DATABASE_USER --> user for the database where the employees should be stored
* ENV DATABASE_PASSWORD --> password for the user mentioned above
* ENV DATABASE_HOST --> host url for the physical database where the employees should be stored
* ENV DATABASE_PORT --> port for the physical database where the employees should be stored
* ENV DATABASE_DB --> database name where the employees should be stored
* ENV PORT --> exposed API port

If the specific environment variable is not set, the api loads the missing values fom config.yaml file. So it is also possible to check in e.g. all variables in config.yaml except the DATABASE_PASSWORD. This can be set via env vars instead.

## example docker build command

docker build --build-arg GO_PATH_ARG="/usr/local/go/bin" -t api_docker_master_kurs:v1.0 -f Dockerfile.fresh .
