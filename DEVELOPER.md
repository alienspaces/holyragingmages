# Developer

## Prerequisites

* Go 1.13
  * [https://golang.org/](https://golang.org/)
* gvm
  * [https://github.com/moovweb/gvm](https://github.com/moovweb/gvm)
* Docker
  * [https://docs.docker.com/install/](https://docs.docker.com/install/)
* psql (Postgresql client)
  * [https://www.postgresql.org/download/macosx/](https://www.postgresql.org/download/macosx/)
  * [https://www.postgresql.org/download/linux/ubuntu/](https://www.postgresql.org/download/linux/ubuntu/)

## Database Migrations

```bash
RUN go get -u -d github.com/golang-migrate/migrate/cmd/migrate github.com/lib/pq github.com/hashicorp/go-multierror \
    && go build -tags 'postgres' -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/cmd/migrate
```

Create new migration

```bash
./script/db-migrate-create [service] [description]
```

Migrate up

```bash
./script/db-migrate-up
```

Migrate down

```bash
./script/db-migrate-down
```

## Testing

Test all packages

```bash
./script/test
```
