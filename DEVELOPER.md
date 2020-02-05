# Developer

## Prerequisites

* Go 1.13
* gvm
* Docker
* psql (postgres client)

## Database Migrations

```bash
RUN go get -u -d github.com/golang-migrate/migrate/cmd/migrate github.com/lib/pq github.com/hashicorp/go-multierror \
    && go build -tags 'postgres' -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/cmd/migrate
```

Create new migration

```bash
./script/db-migrate-create [service [description]
```

Migrate uo

```bash
./script/db-migrate-up
```

Migrate down

```bash
./script/db-migrate-down
```

## Run

Run all services

```bash
./script/run
```
