# Performance Test of Example Go / Node.js CRUD API with SQLite or Postgres

## Performance Test Results: Node.js REST API with SQLite

```sh
TEST_PARALLEL=1 TEST_LIMIT=500 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-21T05:45:42.119Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":500,"TEST_PARALLEL":1,"N_BATCHES":500,"testCount":{"error":0,"success":500,"total":500},"testElapsed":{"count":500,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":500,"min":0,"max":13,"avg":0.734,"p90":1,"p95":1,"p99":2},"readElapsed":{"count":1500,"min":0,"max":14,"avg":0.488,"p90":1,"p95":1,"p99":1},"updateElapsed":{"count":500,"min":0,"max":6,"avg":0.94,"p90":1,"p95":2,"p99":2},"deleteElapsed":{"count":500,"min":0,"max":4,"avg":0.572,"p90":1,"p95":1,"p99":1},"requestElapsed":{"count":3000,"min":0,"max":14,"avg":0.6183333333333333,"p90":1,"p95":1,"p99":2},"requests":{"totalCount":3000,"countPerSecond":1594.896331738437},"elapsedTotal":1881}
# {"timestamp":"2025-07-21T05:46:05.740Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":500,"TEST_PARALLEL":1,"N_BATCHES":500,"testCount":{"error":0,"success":500,"total":500},"testElapsed":{"count":500,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":500,"min":0,"max":18,"avg":0.962,"p90":1,"p95":2,"p99":4},"readElapsed":{"count":1500,"min":0,"max":24,"avg":0.6393333333333333,"p90":1,"p95":1,"p99":3},"updateElapsed":{"count":500,"min":0,"max":8,"avg":1.226,"p90":2,"p95":2,"p99":5},"deleteElapsed":{"count":500,"min":0,"max":13,"avg":0.774,"p90":1,"p95":1,"p99":5},"requestElapsed":{"count":3000,"min":0,"max":24,"avg":0.8133333333333334,"p90":1,"p95":2,"p99":4},"requests":{"totalCount":3000,"countPerSecond":1212.121212121212},"elapsedTotal":2475}
# {"timestamp":"2025-07-21T05:46:18.882Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":500,"TEST_PARALLEL":1,"N_BATCHES":500,"testCount":{"error":0,"success":500,"total":500},"testElapsed":{"count":500,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":500,"min":0,"max":12,"avg":0.68,"p90":1,"p95":2,"p99":3},"readElapsed":{"count":1500,"min":0,"max":6,"avg":0.422,"p90":1,"p95":1,"p99":2},"updateElapsed":{"count":500,"min":0,"max":4,"avg":0.736,"p90":1,"p95":2,"p99":3},"deleteElapsed":{"count":500,"min":0,"max":7,"avg":0.496,"p90":1,"p95":1,"p99":2},"requestElapsed":{"count":3000,"min":0,"max":12,"avg":0.5296666666666666,"p90":1,"p95":1,"p99":2},"requests":{"totalCount":3000,"countPerSecond":1852.9956763434218},"elapsedTotal":1619}

TEST_PARALLEL=100 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-20T15:53:58.416Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":100,"N_BATCHES":100,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":4,"max":96,"avg":10.7499,"p90":13,"p95":14,"p99":27},"readElapsed":{"count":30000,"min":0,"max":31,"avg":6.535666666666667,"p90":10,"p95":11,"p99":16},"updateElapsed":{"count":10000,"min":1,"max":30,"avg":8.8136,"p90":11,"p95":13,"p99":18},"deleteElapsed":{"count":10000,"min":0,"max":22,"avg":6.7857,"p90":10,"p95":11,"p99":15},"requestElapsed":{"count":60000,"min":0,"max":96,"avg":7.659366666666667,"p90":11,"p95":13,"p99":17},"requests":{"totalCount":60000,"countPerSecond":11773.940345368917},"elapsedTotal":5096}
# {"timestamp":"2025-07-20T15:54:33.127Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":100,"N_BATCHES":100,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":4,"max":46,"avg":10.4986,"p90":13,"p95":14,"p99":31},"readElapsed":{"count":30000,"min":0,"max":36,"avg":6.510366666666667,"p90":10,"p95":11,"p99":15},"updateElapsed":{"count":10000,"min":3,"max":34,"avg":8.7311,"p90":11,"p95":13,"p99":20},"deleteElapsed":{"count":10000,"min":0,"max":24,"avg":6.8687,"p90":10,"p95":11,"p99":15},"requestElapsed":{"count":60000,"min":0,"max":46,"avg":7.604916666666667,"p90":11,"p95":12,"p99":17},"requests":{"totalCount":60000,"countPerSecond":11871.784724970319},"elapsedTotal":5054}
# {"timestamp":"2025-07-20T15:54:44.747Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":100,"N_BATCHES":100,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":4,"max":48,"avg":10.5313,"p90":13,"p95":14,"p99":27},"readElapsed":{"count":30000,"min":0,"max":35,"avg":6.446733333333333,"p90":10,"p95":11,"p99":16},"updateElapsed":{"count":10000,"min":2,"max":57,"avg":8.7778,"p90":11,"p95":13,"p99":19},"deleteElapsed":{"count":10000,"min":0,"max":23,"avg":6.8236,"p90":10,"p95":11,"p99":13},"requestElapsed":{"count":60000,"min":0,"max":57,"avg":7.5788166666666665,"p90":11,"p95":12,"p99":17},"requests":{"totalCount":60000,"countPerSecond":11811.023622047243},"elapsedTotal":5080}

TEST_PARALLEL=500 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-20T15:54:55.723Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":500,"N_BATCHES":20,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":24,"max":219,"avg":97.961,"p90":176,"p95":189,"p99":205},"readElapsed":{"count":30000,"min":0,"max":118,"avg":21.0743,"p90":38,"p95":46,"p99":60},"updateElapsed":{"count":10000,"min":2,"max":119,"avg":23.6692,"p90":39,"p95":53,"p99":67},"deleteElapsed":{"count":10000,"min":0,"max":51,"avg":18.1051,"p90":30,"p95":34,"p99":49},"requestElapsed":{"count":60000,"min":0,"max":219,"avg":33.826366666666665,"p90":59,"p95":143,"p99":186},"requests":{"totalCount":60000,"countPerSecond":11204.481792717086},"elapsedTotal":5355}

TEST_PARALLEL=1000 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-20T15:55:14.228Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":1000,"N_BATCHES":10,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":52,"max":510,"avg":206.7932,"p90":400,"p95":446,"p99":488},"readElapsed":{"count":30000,"min":0,"max":200,"avg":35.62263333333333,"p90":63,"p95":77,"p99":117},"updateElapsed":{"count":10000,"min":0,"max":184,"avg":52.3958,"p90":99,"p95":117,"p99":127},"deleteElapsed":{"count":10000,"min":0,"max":92,"avg":32.9476,"p90":58,"p95":84,"p99":87},"requestElapsed":{"count":60000,"min":0,"max":510,"avg":66.50075,"p90":156,"p95":210,"p99":431},"requests":{"totalCount":60000,"countPerSecond":10873.504893077203},"elapsedTotal":5518}
```

## Performance Test Results: Go REST API with SQLite

**NOTE: response times for Go/SQLite are surprisingly slow compared to Go/Postgres and compared to Node.js/SQLite so maybe something is not properly configured?**

```sh
./scripts/performance-test/run.js
# {"timestamp":"2025-07-21T05:41:33.520Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":100,"N_BATCHES":100,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":2,"max":1098,"avg":57.7475,"p90":123,"p95":161,"p99":474},"readElapsed":{"count":30000,"min":0,"max":59,"avg":1.1487,"p90":3,"p95":5,"p99":10},"updateElapsed":{"count":10000,"min":0,"max":989,"avg":10.937,"p90":26,"p95":59,"p99":125},"deleteElapsed":{"count":10000,"min":0,"max":266,"avg":2.782,"p90":4,"p95":10,"p99":40},"requestElapsed":{"count":60000,"min":0,"max":1098,"avg":12.485433333333333,"p90":34,"p95":66,"p99":158},"requests":{"totalCount":60000,"countPerSecond":1662.8789978382572},"elapsedTotal":36082}
```

## Performance Test Results: Go REST API with Postgres

```sh
./scripts/performance-test/run.js
# {"timestamp":"2025-07-21T05:37:01.532Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":100,"N_BATCHES":100,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":2,"max":162,"avg":10.5691,"p90":16,"p95":18,"p99":25},"readElapsed":{"count":30000,"min":0,"max":117,"avg":6.455666666666667,"p90":9,"p95":11,"p99":15},"updateElapsed":{"count":10000,"min":1,"max":83,"avg":8.2816,"p90":11,"p95":13,"p99":17},"deleteElapsed":{"count":10000,"min":0,"max":19,"avg":6.3088,"p90":8,"p95":9,"p99":12},"requestElapsed":{"count":60000,"min":0,"max":162,"avg":7.421083333333334,"p90":11,"p95":14,"p99":18},"requests":{"totalCount":60000,"countPerSecond":11848.341232227487},"elapsedTotal":5064}
# {"timestamp":"2025-07-21T05:39:01.113Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":100,"N_BATCHES":100,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":2,"max":50,"avg":9.5874,"p90":15,"p95":17,"p99":26},"readElapsed":{"count":30000,"min":0,"max":37,"avg":5.8463666666666665,"p90":8,"p95":10,"p99":14},"updateElapsed":{"count":10000,"min":1,"max":26,"avg":7.6027,"p90":10,"p95":13,"p99":16},"deleteElapsed":{"count":10000,"min":0,"max":16,"avg":5.7658,"p90":8,"p95":8,"p99":11},"requestElapsed":{"count":60000,"min":0,"max":50,"avg":6.7491666666666665,"p90":10,"p95":13,"p99":18},"requests":{"totalCount":60000,"countPerSecond":13020.833333333334},"elapsedTotal":4608}

TEST_PARALLEL=500 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-21T05:39:48.593Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":500,"N_BATCHES":20,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":23,"max":176,"avg":83.9441,"p90":145,"p95":153,"p99":169},"readElapsed":{"count":30000,"min":0,"max":91,"avg":25.7337,"p90":37,"p95":42,"p99":64},"updateElapsed":{"count":10000,"min":12,"max":124,"avg":28.8022,"p90":37,"p95":56,"p99":66},"deleteElapsed":{"count":10000,"min":0,"max":59,"avg":25.7211,"p90":40,"p95":45,"p99":57},"requestElapsed":{"count":60000,"min":0,"max":176,"avg":35.94475,"p90":62,"p95":117,"p99":152},"requests":{"totalCount":60000,"countPerSecond":11322.891111530476},"elapsedTotal":5299}

TEST_PARALLEL=1000 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-21T05:40:29.742Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":1000,"N_BATCHES":10,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":52,"max":339,"avg":170.1865,"p90":274,"p95":285,"p99":330},"readElapsed":{"count":30000,"min":1,"max":236,"avg":48.630066666666664,"p90":70,"p95":92,"p99":107},"updateElapsed":{"count":10000,"min":25,"max":120,"avg":55.7312,"p90":81,"p95":108,"p99":116},"deleteElapsed":{"count":10000,"min":5,"max":85,"avg":48.7948,"p90":71,"p95":83,"p99":84},"requestElapsed":{"count":60000,"min":1,"max":339,"avg":70.10045,"p90":162,"p95":187,"p99":276},"requests":{"totalCount":60000,"countPerSecond":11547.344110854505},"elapsedTotal":5196}
```

## Performance Test Results: Ruby on Rails REST API with SQLite

```sh
TEST_PARALLEL=10 TEST_LIMIT=100 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-20T21:44:17.729Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":100,"TEST_PARALLEL":10,"N_BATCHES":10,"testCount":{"error":0,"success":100,"total":100},"testElapsed":{"count":100,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":100,"min":20,"max":1900,"avg":541.46,"p90":1229,"p95":1275,"p99":1900},"readElapsed":{"count":300,"min":3,"max":322,"avg":70.78333333333333,"p90":213,"p95":286,"p99":314},"updateElapsed":{"count":100,"min":4,"max":201,"avg":28.13,"p90":50,"p95":134,"p99":201},"deleteElapsed":{"count":100,"min":4,"max":1685,"avg":153.47,"p90":1162,"p95":1260,"p99":1685},"requestElapsed":{"count":600,"min":3,"max":1900,"avg":155.90166666666667,"p90":555,"p95":1025,"p99":1300},"requests":{"totalCount":600,"countPerSecond":38.33375926399182},"elapsedTotal":15652}

TEST_PARALLEL=20 TEST_LIMIT=200 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-20T21:47:25.674Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":200,"TEST_PARALLEL":20,"N_BATCHES":10,"testCount":{"error":0,"success":200,"total":200},"testElapsed":{"count":200,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":200,"min":20,"max":3009,"avg":1287.975,"p90":2386,"p95":2516,"p99":2883},"readElapsed":{"count":600,"min":3,"max":407,"avg":52.15,"p90":118,"p95":205,"p99":261},"updateElapsed":{"count":200,"min":4,"max":239,"avg":18.26,"p90":29,"p95":107,"p99":238},"deleteElapsed":{"count":200,"min":4,"max":2654,"avg":137.35,"p90":23,"p95":2438,"p99":2645},"requestElapsed":{"count":1200,"min":3,"max":3009,"avg":266.6725,"p90":1220,"p95":2007,"p99":2580},"requests":{"totalCount":1200,"countPerSecond":42.67273567796309},"elapsedTotal":28121}

TEST_PARALLEL=30 TEST_LIMIT=300 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-21T05:31:15.749Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":300,"TEST_PARALLEL":30,"N_BATCHES":10,"testCount":{"error":0,"success":300,"total":300},"testElapsed":{"count":300,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":300,"min":21,"max":4056,"avg":1878.3633333333332,"p90":3447,"p95":3617,"p99":3810},"readElapsed":{"count":900,"min":3,"max":337,"avg":45.15,"p90":109,"p95":121,"p99":247},"updateElapsed":{"count":300,"min":4,"max":229,"avg":13.47,"p90":23,"p95":32,"p99":199},"deleteElapsed":{"count":300,"min":4,"max":3895,"avg":134.23666666666668,"p90":21,"p95":107,"p99":3805},"requestElapsed":{"count":1800,"min":3,"max":4056,"avg":360.25333333333333,"p90":1678,"p95":2808,"p99":3699},"requests":{"totalCount":1800,"countPerSecond":45.24886877828054},"elapsedTotal":39780}

TEST_PARALLEL=50 TEST_LIMIT=150 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-21T05:32:16.644Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":150,"TEST_PARALLEL":50,"N_BATCHES":3,"testCount":{"error":0,"success":150,"total":150},"testElapsed":{"count":150,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":150,"min":24,"max":7154,"avg":3553.806666666667,"p90":6150,"p95":6628,"p99":7027},"readElapsed":{"count":450,"min":3,"max":371,"avg":49.89333333333333,"p90":112,"p95":139,"p99":344},"updateElapsed":{"count":150,"min":4,"max":129,"avg":10.373333333333333,"p90":23,"p95":25,"p99":126},"deleteElapsed":{"count":150,"min":4,"max":6242,"avg":91.58,"p90":21,"p95":41,"p99":6142},"requestElapsed":{"count":900,"min":3,"max":7154,"avg":634.24,"p90":3132,"p95":5016,"p99":6531},"requests":{"totalCount":900,"countPerSecond":43.30045706038008},"elapsedTotal":20785}

TEST_PARALLEL=80 TEST_LIMIT=400 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-21T05:34:07.439Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":400,"TEST_PARALLEL":80,"N_BATCHES":5,"testCount":{"error":0,"success":400,"total":400},"testElapsed":{"count":400,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":400,"min":26,"max":11802,"avg":5288.0325,"p90":9346,"p95":10149,"p99":11428},"readElapsed":{"count":1200,"min":3,"max":360,"avg":43.87833333333333,"p90":110,"p95":120,"p99":306},"updateElapsed":{"count":400,"min":4,"max":119,"avg":8.6625,"p90":16,"p95":23,"p99":108},"deleteElapsed":{"count":400,"min":4,"max":10430,"avg":130.26,"p90":13,"p95":22,"p99":9423},"requestElapsed":{"count":2400,"min":3,"max":11802,"avg":926.4316666666666,"p90":4429,"p95":7446,"p99":10027},"requests":{"totalCount":2400,"countPerSecond":45.314653626116346},"elapsedTotal":52963}
```

## Developer Setup - Go Server

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

Command to stop any server started somewhere else (not needed initially)

```sh
scripts/stop
```

## Developer Setup - Node.js Server

```sh
nvm use # Node 24

npm install

npm run dev
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

## How the Ruby on Rails API Was Created

```sh
# Install Ruby and Ruby on Rails
# https://guides.rubyonrails.org/install_ruby_on_rails.html#install-ruby-on-macos
mise use -g ruby@3
ruby --version
# ruby 3.4.5 (2025-07-16 revision 20cda200d3) +PRISM [arm64-darwin24]
gem install rails
# Successfully installed rails-8.0.2
# https://guides.rubyonrails.org/getting_started.html
# In new terminal:
rails --version
gem update

# Create Rails API
rails new --api content-api-ruby
cd content-api-ruby
bin/rails server
bin/rails g scaffold content id:text title:text body:text author:text status:text data:text --api
#  invoke  active_record
#       remove    db/migrate/20250720202444_create_contents.rb
#       create    db/migrate/20250720205608_create_contents.rb
#    identical    app/models/content.rb
#       invoke    test_unit
#    identical      test/models/content_test.rb
#    identical      test/fixtures/contents.yml
#       invoke  resource_route
#        route    resources :contents
#       invoke  scaffold_controller
#        force    app/controllers/contents_controller.rb
#       invoke    resource_route
#       invoke    test_unit
#        force      test/controllers/contents_controller_test.rb

bin/rails db:migrate
bin/rails server -p 8888

# EDITS:
# edit migration to have same SQL as the Go/Node.js app but with table name contents
# edit routes.rb to have: resources :contents, path: 'content'
# edit contents_controller.rb content_params to return:
#   {
#     title: params[:title],
#     body: params[:body],
#     author: params[:author],
#     status: params[:status],
#     data: params[:data]
#   }
```

## Resources

- [go-sqlite3 - SQLite Library](https://github.com/mattn/go-sqlite3)
- [pgx - Postgres Library](https://github.com/jackc/pgx)
- [WAL File Size Issue with SQLite](https://news.ycombinator.com/item?id=40688987)

- [Huma Web Framework with OpenAPI Support](https://github.com/danielgtaylor/huma)
- [Huma Logging Middleware with Request Context](https://github.com/danielgtaylor/huma/blob/v1.14.3/middleware/logger.go)
