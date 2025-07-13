# Example Go CRUD API with SQLite/Postgres and OpenAPI

## Developer Setup

```sh
# Check go installation
go version
# go version go1.24.4 darwin/arm64

# Install dependencies
go get

# Install migrate CLI and run migrations
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
scripts/migrate

# Start server
scripts/run

# Run smoke test
scripts/smoke-test.sh

# Load SQLite test data
go run scripts/sqlite/generate_test_data/main.go 
```

Pretty printing JSON logs:

```sh
./scripts/run | tee log.json

# In other terminal
tail -f log.json | jq
```

Useful command to stop any server running somewhere else (not needed initially)

```sh
scripts/stop
```

## Invoking the API with Curl

```sh
scripts/smoke-test.sh
```

## Generate Test Data

```sh
rm -f db/sqlite/content-api.db
./scripts/migrate

# Without WAL mode:
go run scripts/sqlite/generate_test_data/main.go 
# Successfully created 100000 test records in 46.16 seconds
# Average rate: 2166.5 records/second
du -sh db/sqlite/content-api.db 
# 61M	db/sqlite/content-api.db

# Without WAL mode:
go run scripts/sqlite/generate_test_data/main.go 
# Successfully created 100000 test records in 8.57 seconds
# Average rate: 11664.8 records/second
```

## Running the Server with Postgres

```sh
# Assuming you have postgres running locally with Postgres.app etc.
createdb content_api

# Running the migration
migrate -database "postgres://postgres:postgres@localhost:5432/content_api?sslmode=disable" -path db/postgres/migrations up

# Start server and make it talk to postgres
DATABASE_ENGINE=postgres ./scripts/run

# Test
./scripts/smoke-test.sh
```


## SQLite Performance Test

```sh
go run scripts/sqlite/performance-test/main.go -n 1000 -parallel 1
# 2025/07/11 11:37:26 DELETE: 1000 total, 1000 success, 0 failures, avg: 468.915µs, min: 213.375µs, max: 14.076375ms
# 2025/07/11 11:37:26 CREATE: 1000 total, 1000 success, 0 failures, avg: 175.512µs, min: 104.042µs, max: 3.078ms
# 2025/07/11 11:37:26 READ: 1000 total, 1000 success, 0 failures, avg: 398.843µs, min: 127.917µs, max: 13.551709ms
# 2025/07/11 11:37:26 UPDATE: 1000 total, 1000 success, 0 failures, avg: 474.764µs, min: 333.667µs, max: 3.897708ms
# 2025/07/11 11:37:26 
# Total: 4000 operations, 4000 success, 0 failures

go run scripts/sqlite/performance-test/main.go -n 1000 -parallel 10
# 2025/07/11 11:37:50 CREATE: 1000 total, 1000 success, 0 failures, avg: 1.468545ms, min: 104µs, max: 109.819416ms
# 2025/07/11 11:37:50 READ: 1000 total, 1000 success, 0 failures, avg: 710.899µs, min: 152.417µs, max: 2.082917ms
# 2025/07/11 11:37:50 UPDATE: 1000 total, 1000 success, 0 failures, avg: 2.253803ms, min: 254.209µs, max: 140.495292ms
# 2025/07/11 11:37:50 DELETE: 1000 total, 1000 success, 0 failures, avg: 1.74087ms, min: 207.291µs, max: 83.590333ms
# 2025/07/11 11:37:50 
# Total: 4000 operations, 4000 success, 0 failures

go run scripts/sqlite/performance-test/main.go -n 10000 -parallel 50
# 2025/07/11 17:31:58 CREATE: 10000 total, 10000 success, 0 failures, avg: 7.274489ms, min: 101.709µs, max: 1.363181834s
# 2025/07/11 17:31:58 READ: 10000 total, 10000 success, 0 failures, avg: 2.937146ms, min: 126µs, max: 13.51125ms
# 2025/07/11 17:31:58 UPDATE: 10000 total, 10000 success, 0 failures, avg: 20.08348ms, min: 248.708µs, max: 876.129125ms
# 2025/07/11 17:31:58 DELETE: 10000 total, 10000 success, 0 failures, avg: 47.864971ms, min: 3.0175ms, max: 293.351917ms
# 2025/07/11 17:31:58 
# Total: 40000 operations, 40000 success, 0 failures

go run scripts/sqlite/performance-test/main.go -n 10000 -parallel 100
# 2025/07/11 17:34:56 DELETE: 10000 total, 10000 success, 0 failures, avg: 96.714906ms, min: 15.95825ms, max: 290.583084ms
# 2025/07/11 17:34:56 CREATE: 10000 total, 10000 success, 0 failures, avg: 12.601139ms, min: 99.833µs, max: 1.468209292s
# 2025/07/11 17:34:56 READ: 10000 total, 10000 success, 0 failures, avg: 6.335471ms, min: 127.083µs, max: 43.939583ms
# 2025/07/11 17:34:56 UPDATE: 10000 total, 10000 success, 0 failures, avg: 40.445279ms, min: 278.792µs, max: 1.183335042s
# 2025/07/11 17:34:56 
# Total: 40000 operations, 40000 success, 0 failures
# du -sh db/sqlite/*                                          
#  98M	db/sqlite/content-api.db
#  32K	db/sqlite/content-api.db-shm
#  6.9M	db/sqlite/content-api.db-wal
```

## Postgres Performance Test

```sh
go run scripts/sqlite/performance-test/main.go -n 1000 -parallel 1
# 2025/07/11 17:17:29 DELETE: 1000 total, 1000 success, 0 failures, avg: 400.562µs, min: 253.666µs, max: 23.594417ms
# 2025/07/11 17:17:29 CREATE: 1000 total, 1000 success, 0 failures, avg: 252.243µs, min: 143.75µs, max: 7.908916ms
# 2025/07/11 17:17:29 READ: 1000 total, 1000 success, 0 failures, avg: 419.485µs, min: 309µs, max: 1.727709ms
# 2025/07/11 17:17:29 UPDATE: 1000 total, 1000 success, 0 failures, avg: 631.14µs, min: 406.042µs, max: 3.596208ms
# 2025/07/11 17:17:29 
# Total: 4000 operations, 4000 success, 0 failures

go run scripts/sqlite/performance-test/main.go -n 1000 -parallel 10
# 2025/07/11 17:18:00 CREATE: 1000 total, 1000 success, 0 failures, avg: 1.172596ms, min: 288.125µs, max: 47.014625ms
# 2025/07/11 17:18:00 READ: 1000 total, 1000 success, 0 failures, avg: 944.524µs, min: 360.292µs, max: 2.751167ms
# 2025/07/11 17:18:00 UPDATE: 1000 total, 1000 success, 0 failures, avg: 1.466389ms, min: 692.625µs, max: 4.718209ms
# 2025/07/11 17:18:00 DELETE: 1000 total, 1000 success, 0 failures, avg: 980.447µs, min: 435.917µs, max: 2.570375ms
# 2025/07/11 17:18:00 
# Total: 4000 operations, 4000 success, 0 failures

go run scripts/sqlite/performance-test/main.go -n 10000 -parallel 50
# 2025/07/11 17:30:56 CREATE: 10000 total, 10000 success, 0 failures, avg: 3.9116ms, min: 273.209µs, max: 156.371416ms
# 2025/07/11 17:30:56 READ: 10000 total, 10000 success, 0 failures, avg: 4.196421ms, min: 850.833µs, max: 31.755292ms
# 2025/07/11 17:30:56 UPDATE: 10000 total, 10000 success, 0 failures, avg: 34.095279ms, min: 1.63075ms, max: 123.662667ms
# 2025/07/11 17:30:56 DELETE: 10000 total, 9691 success, 309 failures, avg: 57.581314ms, min: 6.896917ms, max: 1.130187916s
# 2025/07/11 17:30:56 
# Total: 40000 operations, 39691 success, 309 failures
```

## API Docs and OpenAPI Specification

```sh
# API Docs
open http://localhost:8888/docs

# OpenAPI Specification
curl -s http://localhost:8888/openapi.yaml | yq
```

## SQLite Database Migrations

Using the `golang-migrate` package:

```sh
# Install golang-migrate with SQLite support
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

```sh
# Creating a migration
migrate create -ext sql -dir db/sqlite/migrations -seq create_content_table
# db/sqlite/migrations/000001_create_content_table.up.sql
# db/sqlite/migrations/000001_create_content_table.down.sql

# Running a migration
migrate -database "sqlite3://db/sqlite/content-api.db" -path db/sqlite/migrations up
```

## Postgres Database Migrations

Using the `golang-migrate` package:

```sh
# Creating a migration
migrate create -ext sql -dir db/postgres/migrations -seq create_content_table

# Assuming you have postgres running locally with Postgres.app etc.
createdb content_api

# Running the migration
migrate -database "postgres://postgres:postgres@localhost:5432/content_api?sslmode=disable" -path db/postgres/migrations up
```

# SQLite Console

```sh
sqlite3 db/sqlite/content-api.db 
.schema
select count(*) from content;
select * from content order by created_at desc limit 100;
```

## Resources

* [go-sqlite3 - SQLite Library](https://github.com/mattn/go-sqlite3)
* [pgx - Postgres Library](https://github.com/jackc/pgx)
* [WAL File Size Issue with SQLite](https://news.ycombinator.com/item?id=40688987)

* [Huma Web Framework with OpenAPI Support](https://github.com/danielgtaylor/huma)
* [Huma Logging Middleware with Request Context](https://github.com/danielgtaylor/huma/blob/v1.14.3/middleware/logger.go)

