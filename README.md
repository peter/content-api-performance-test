# Performance Test of of Single Server CRUD/REST API with SQLite or Postgres (Using Node.js / Go / Elixir / Ruby)

The motivation for these performance tests was to figure out if you can get sufficient performance with SQLite running in a single process on a single server. The use case I had in mind was a CMS backend (or similar type of app) exposing a classic CRUD/REST API.

**Preliminary finding: on my M1 Macbook laptop Node.js with either SQLite or Postgres can do around 10k read and write queries per second with a p99 response time of around 20 ms and request concurrency level of 100. This should be more than sufficient for most applications.**

## Performance Test Results Summary

The [test runner](scripts/performance-test/run.js) is written in JavaScript and issues HTTP requests against a web server running on localhost on port 8888. The script will execute a certain number of tests in parallel (i.e. 100) and each test executes the following sequence of HTTP requests:

1. POST /content
2. GET /content/:id
3. PUT /content/:id
4. GET /content/:id
5. DELETE /content/:id
6. GET /content/:id

The content being created is stored in a `content` database table and looks roughly like this:

```javascript
{
    title:  `Smoke Test Content ${runId}`,
    body:   `This is smoke test content number ${runId}`,
    author: "Smoke Tester",
    status: "draft",
    data: {
        "run_id": runId,
        "created_at": createdAt,
    },
}
```

The numbers below should be considered rough estimations (ballpark figures) and no particular optimization has been made in each runtime to improve performance or scalability. The code between each platform can differ slightly (in how it validates payloads etc.):

|Runtime|Database|Number of Tests|Test Concurrency|Number of Requests|Total Time (ms)|Requests Per Second|Response Time Avg (ms)|Response Time P99 (ms)|Read Response Time Avg (ms)|Read Response Time P99 (ms)|
|-|-|-|-|-|-|-|-|-|-|-|
|Node.js|SQLite|500|1|3000|950|3200|0.3|1|0.3|1|
|Node.js|SQLite|10000|10|60000|5300|11000|0.8|2|0.7|2|
|Node.js|SQLite|10000|100|60000|5000|12000|7|17|6|16|
|Node.js|SQLite|10000|1000|60000|6100|9900|57|464|20|152|
|Node.js|Memory|500|1|3000|825|3600|0.3|1|0.2|1|
|Node.js|Memory|10000|100|60000|4375|14000|6|15|6|15|
|Node.js|Postgres|500|1|3000|1900|1600|0.6|2|0.5|1|
|Node.js|Postgres|10000|100|60000|5100|12000|7|17|7|16|
|Node.js|Postgres|10000|1000|60000|5500|11000|67|431|36|117|
|Go|SQLite|10000|100|60000|36000|1700|12|158|1|10|
|Go|Postgres|10000|100|60000|5100|11800|7|18|6|15|
|Elixir/Phoenix|SQLite|10000|100|60000|29000|2100|42|104|48|112|
|Elixir/Phoenix|Postgres|10000|100|60000|29900|2004|43|107|51|118|
|Ruby on Rails|SQLite|400|80|2400|30000|79|500|5700|4|13|

## Performance Test Results: Node.js REST API with SQLite

* Repo: https://github.com/peter/content-api-performance-test
* Start server command: `npm run dev`

```sh
TEST_PARALLEL=1 TEST_LIMIT=500 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-23T06:23:29.885Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":500,"TEST_PARALLEL":1,"N_BATCHES":500,"testCount":{"error":0,"success":500,"total":500},"testElapsed":{"count":500,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":500,"min":0,"max":12,"avg":0.4,"p90":1,"p95":1,"p99":1},"readElapsed":{"count":1500,"min":0,"max":2,"avg":0.25266666666666665,"p90":1,"p95":1,"p99":1},"updateElapsed":{"count":500,"min":0,"max":3,"avg":0.394,"p90":1,"p95":1,"p99":1},"deleteElapsed":{"count":500,"min":0,"max":2,"avg":0.31,"p90":1,"p95":1,"p99":1},"requestElapsed":{"count":3000,"min":0,"max":12,"avg":0.31033333333333335,"p90":1,"p95":1,"p99":1},"requests":{"totalCount":3000,"countPerSecond":3151.2605042016808},"elapsedTotal":952}

TEST_PARALLEL=10 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-25T17:50:34.750Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":10,"N_BATCHES":1000,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":0,"max":27,"avg":1.2225,"p90":2,"p95":2,"p99":3},"readElapsed":{"count":30000,"min":0,"max":5,"avg":0.6598333333333334,"p90":1,"p95":2,"p99":2},"updateElapsed":{"count":10000,"min":0,"max":5,"avg":0.8007,"p90":1,"p95":2,"p99":2},"deleteElapsed":{"count":10000,"min":0,"max":5,"avg":0.5473,"p90":1,"p95":1,"p99":2},"requestElapsed":{"count":60000,"min":0,"max":27,"avg":0.7583333333333333,"p90":1,"p95":2,"p99":2},"requests":{"totalCount":60000,"countPerSecond":11355.034065102196},"elapsedTotal":5284}

TEST_PARALLEL=100 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-23T06:21:42.687Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":100,"N_BATCHES":100,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":3,"max":102,"avg":9.3115,"p90":14,"p95":15,"p99":29},"readElapsed":{"count":30000,"min":0,"max":31,"avg":6.399533333333333,"p90":11,"p95":12,"p99":16},"updateElapsed":{"count":10000,"min":1,"max":26,"avg":9.0198,"p90":12,"p95":14,"p99":18},"deleteElapsed":{"count":10000,"min":0,"max":18,"avg":5.1821,"p90":8,"p95":9,"p99":11},"requestElapsed":{"count":60000,"min":0,"max":102,"avg":7.118666666666667,"p90":12,"p95":13,"p99":17},"requests":{"totalCount":60000,"countPerSecond":11961.722488038278},"elapsedTotal":5016}

TEST_PARALLEL=500 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-23T06:21:57.712Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":500,"N_BATCHES":20,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":24,"max":259,"avg":106.043,"p90":207,"p95":221,"p99":236},"readElapsed":{"count":30000,"min":0,"max":121,"avg":13.410566666666666,"p90":30,"p95":41,"p99":77},"updateElapsed":{"count":10000,"min":0,"max":116,"avg":15.2661,"p90":30,"p95":38,"p99":60},"deleteElapsed":{"count":10000,"min":0,"max":59,"avg":10.3394,"p90":20,"p95":33,"p99":49},"requestElapsed":{"count":60000,"min":0,"max":259,"avg":28.6467,"p90":56,"p95":155,"p99":218},"requests":{"totalCount":60000,"countPerSecond":11398.176291793312},"elapsedTotal":5264}

TEST_PARALLEL=1000 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-23T06:22:10.599Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":1000,"N_BATCHES":10,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":58,"max":740,"avg":241.3515,"p90":446,"p95":470,"p99":502},"readElapsed":{"count":30000,"min":0,"max":193,"avg":20.398033333333334,"p90":55,"p95":71,"p99":152},"updateElapsed":{"count":10000,"min":0,"max":162,"avg":25.3618,"p90":58,"p95":113,"p99":158},"deleteElapsed":{"count":10000,"min":0,"max":81,"avg":16.5591,"p90":40,"p95":67,"p99":77},"requestElapsed":{"count":60000,"min":0,"max":740,"avg":57.41108333333333,"p90":167,"p95":326,"p99":464},"requests":{"totalCount":60000,"countPerSecond":9894.459102902374},"elapsedTotal":6064}
```

## Performance Test Results: Node.js REST API with Memory Store

* Repo: https://github.com/peter/content-api-performance-test
* Start server command: `DATABASE_ENGINE=memory npm run dev`

```sh
TEST_PARALLEL=1 TEST_LIMIT=500 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-25T17:44:31.190Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":500,"TEST_PARALLEL":1,"N_BATCHES":500,"testCount":{"error":0,"success":500,"total":500},"testElapsed":{"count":500,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":500,"min":0,"max":13,"avg":0.378,"p90":1,"p95":1,"p99":1},"readElapsed":{"count":1500,"min":0,"max":2,"avg":0.24933333333333332,"p90":1,"p95":1,"p99":1},"updateElapsed":{"count":500,"min":0,"max":2,"avg":0.278,"p90":1,"p95":1,"p99":1},"deleteElapsed":{"count":500,"min":0,"max":1,"avg":0.184,"p90":1,"p95":1,"p99":1},"requestElapsed":{"count":3000,"min":0,"max":13,"avg":0.26466666666666666,"p90":1,"p95":1,"p99":1},"requests":{"totalCount":3000,"countPerSecond":3636.3636363636365},"elapsedTotal":825}

TEST_PARALLEL=100 ./scripts/performance-test/run.js
# {"timestamp":"2025-07-25T17:44:47.227Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":100,"N_BATCHES":100,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":3,"max":89,"avg":7.3695,"p90":10,"p95":11,"p99":29},"readElapsed":{"count":30000,"min":0,"max":31,"avg":6.172066666666667,"p90":11,"p95":12,"p99":15},"updateElapsed":{"count":10000,"min":1,"max":25,"avg":6.6081,"p90":9,"p95":12,"p99":15},"deleteElapsed":{"count":10000,"min":0,"max":17,"avg":5.8797,"p90":8,"p95":9,"p99":11},"requestElapsed":{"count":60000,"min":0,"max":89,"avg":6.395583333333334,"p90":10,"p95":12,"p99":15},"requests":{"totalCount":60000,"countPerSecond":13770.943309616709},"elapsedTotal":4357}
```

## Performance Test Results: Node.js REST API with Postgres

* Repo: https://github.com/peter/content-api-performance-test
* Start server command: `DATABASE_ENGINE=postgres npm run dev`

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

* Repo: https://github.com/peter/content-api-ruby
* Start server command: `RAILS_ENV=production bin/rails server -p 8888`

```sh
TEST_PARALLEL=10 TEST_LIMIT=100 TEST_DATA_FIELD=false ./scripts/performance-test/run.js
# {"timestamp":"2025-07-22T15:41:57.406Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":100,"TEST_PARALLEL":10,"N_BATCHES":10,"testCount":{"error":0,"success":100,"total":100},"testElapsed":{"count":100,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":100,"min":2,"max":755,"avg":255.42,"p90":499,"p95":696,"p99":755},"readElapsed":{"count":300,"min":1,"max":19,"avg":4.913333333333333,"p90":10,"p95":14,"p99":18},"updateElapsed":{"count":100,"min":1,"max":16,"avg":4.26,"p90":8,"p95":12,"p99":16},"deleteElapsed":{"count":100,"min":1,"max":725,"avg":37.38,"p90":6,"p95":679,"p99":725},"requestElapsed":{"count":600,"min":1,"max":755,"avg":51.96666666666667,"p90":239,"p95":477,"p99":700},"requests":{"totalCount":600,"countPerSecond":82.11304228821677},"elapsedTotal":7307}


TEST_PARALLEL=20 TEST_LIMIT=200 TEST_DATA_FIELD=false ./scripts/performance-test/run.js
# {"timestamp":"2025-07-22T15:43:10.976Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":200,"TEST_PARALLEL":20,"N_BATCHES":10,"testCount":{"error":0,"success":200,"total":200},"testElapsed":{"count":200,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":200,"min":2,"max":1486,"avg":604.035,"p90":1187,"p95":1367,"p99":1478},"readElapsed":{"count":600,"min":0,"max":35,"avg":4.336666666666667,"p90":7,"p95":12,"p99":22},"updateElapsed":{"count":200,"min":1,"max":16,"avg":3.305,"p90":5,"p95":7,"p99":13},"deleteElapsed":{"count":200,"min":1,"max":1481,"avg":73.295,"p90":4,"p95":1362,"p99":1478},"requestElapsed":{"count":1200,"min":0,"max":1486,"avg":115.6075,"p90":482,"p95":957,"p99":1394},"requests":{"totalCount":1200,"countPerSecond":83.71703641691084},"elapsedTotal":14334}

TEST_PARALLEL=80 TEST_LIMIT=400 TEST_DATA_FIELD=false ./scripts/performance-test/run.js
# {"timestamp":"2025-07-22T15:43:58.954Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":400,"TEST_PARALLEL":80,"N_BATCHES":5,"testCount":{"error":0,"success":400,"total":400},"testElapsed":{"count":400,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":400,"min":8,"max":6099,"avg":2921.58,"p90":5394,"p95":5706,"p99":6029},"readElapsed":{"count":1200,"min":1,"max":22,"avg":4.161666666666667,"p90":6,"p95":7,"p99":13},"updateElapsed":{"count":400,"min":1,"max":10,"avg":3.1525,"p90":5,"p95":6,"p99":9},"deleteElapsed":{"count":400,"min":1,"max":6070,"avg":62.3675,"p90":3,"p95":4,"p99":5962},"requestElapsed":{"count":2400,"min":1,"max":6099,"avg":499.93083333333334,"p90":2357,"p95":4206,"p99":5706},"requests":{"totalCount":2400,"countPerSecond":79.43074631805395},"elapsedTotal":30215}
```

## Performance Test Result: Elixir/Phoenix with SQLite

* Repo: https://github.com/peter/content-api-elixir
* Start server command: `mix phx.server`

```sh
TEST_PARALLEL=100 TEST_DATA_FIELD=false ./scripts/performance-test/run.js 
# {"timestamp":"2025-07-22T14:27:34.137Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":100,"N_BATCHES":100,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":13,"max":131,"avg":39.697,"p90":57,"p95":66,"p99":103},"readElapsed":{"count":30000,"min":9,"max":138,"avg":48.007533333333335,"p90":87,"p95":94,"p99":112},"updateElapsed":{"count":10000,"min":12,"max":88,"avg":33.6144,"p90":43,"p95":48,"p99":60},"deleteElapsed":{"count":10000,"min":12,"max":82,"avg":34.3503,"p90":44,"p95":50,"p99":63},"requestElapsed":{"count":60000,"min":9,"max":138,"avg":41.947383333333335,"p90":77,"p95":88,"p99":104},"requests":{"totalCount":60000,"countPerSecond":2070.8935905843373},"elapsedTotal":28973}
# {"timestamp":"2025-07-22T14:28:17.237Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":100,"N_BATCHES":100,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":14,"max":88,"avg":38.7593,"p90":55,"p95":61,"p99":71},"readElapsed":{"count":30000,"min":13,"max":131,"avg":48.6483,"p90":90,"p95":96,"p99":109},"updateElapsed":{"count":10000,"min":15,"max":77,"avg":33.6887,"p90":43,"p95":48,"p99":57},"deleteElapsed":{"count":10000,"min":13,"max":80,"avg":33.3294,"p90":42,"p95":46,"p99":56},"requestElapsed":{"count":60000,"min":13,"max":131,"avg":41.953716666666665,"p90":77,"p95":90,"p99":103},"requests":{"totalCount":60000,"countPerSecond":2081.6708878326335},"elapsedTotal":28823}
```

## Performance Test Result: Elixir/Phoenix with Postgres

* Repo: https://github.com/peter/content-api-elixir
* Start server command: `DATABASE_ENGINE=postgres mix phx.server`

```sh
TEST_PARALLEL=100 TEST_DATA_FIELD=false ./scripts/performance-test/run.js 
# {"timestamp":"2025-07-22T17:11:09.893Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":100,"N_BATCHES":100,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":13,"max":112,"avg":39.0961,"p90":54,"p95":62,"p99":78},"readElapsed":{"count":30000,"min":7,"max":182,"avg":51.065,"p90":92,"p95":98,"p99":118},"updateElapsed":{"count":10000,"min":14,"max":96,"avg":34.2748,"p90":43,"p95":49,"p99":68},"deleteElapsed":{"count":10000,"min":13,"max":117,"avg":34.4257,"p90":44,"p95":49,"p99":66},"requestElapsed":{"count":60000,"min":7,"max":182,"avg":43.4986,"p90":82,"p95":92,"p99":107},"requests":{"totalCount":60000,"countPerSecond":2004.4097013429543},"elapsedTotal":29934}
# {"timestamp":"2025-07-22T17:20:28.985Z","level":"INFO","message":"Finished performance test","TEST_LIMIT":10000,"TEST_PARALLEL":100,"N_BATCHES":100,"testCount":{"error":0,"success":10000,"total":10000},"testElapsed":{"count":10000,"min":0,"max":0,"avg":0,"p90":0,"p95":0,"p99":0},"createElapsed":{"count":10000,"min":14,"max":102,"avg":44.2636,"p90":56,"p95":62,"p99":86},"readElapsed":{"count":30000,"min":14,"max":180,"avg":56.072433333333336,"p90":100,"p95":109,"p99":142},"updateElapsed":{"count":10000,"min":15,"max":87,"avg":38.2731,"p90":48,"p95":54,"p99":68},"deleteElapsed":{"count":10000,"min":23,"max":118,"avg":38.3921,"p90":47,"p95":56,"p99":85},"requestElapsed":{"count":60000,"min":14,"max":180,"avg":48.19101666666667,"p90":89,"p95":101,"p99":126},"requests":{"totalCount":60000,"countPerSecond":1858.620903289759},"elapsedTotal":32282}
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

## SQLite Console

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

## How the Elixir App Was Created

```sh
mix phx.new content-api-elixir --app content_api --database sqlite3 --no-html --no-assets
cd content-api-elixir
# mix deps.get
mix ecto.create
# EDIT: Change port in config/dev.exs to 8888
mix phx.server

mix phx.gen.context Api Content content id:text title:text body:text author:text status:text data:text
# * creating lib/content_api/api/content.ex
# * creating priv/repo/migrations/20250722061051_create_content.exs
# * creating lib/content_api/api.ex
# * injecting lib/content_api/api.ex
# * creating test/content_api/api_test.exs
# * injecting test/content_api/api_test.exs
# * creating test/support/fixtures/api_fixtures.ex
# * injecting test/support/fixtures/api_fixtures.ex
# EDIT: add custom SQL to migration file
# EDIT: add @primary_key {:id, :string, []} to model
mix ecto.migrate
mix phx.gen.json Api Content content title:text body:text author:text status:text data:text --no-context --no-schema
# * creating lib/content_api_web/controllers/content_controller.ex
# * creating lib/content_api_web/controllers/content_json.ex
# * creating lib/content_api_web/controllers/changeset_json.ex
# * creating test/content_api_web/controllers/content_controller_test.exs
# * creating lib/content_api_web/controllers/fallback_controller.ex
# EDIT: Add the resource to the "/api" scope in lib/content_api_web/router.ex:
# resources "/content", ContentController, except: [:new, :edit]
# EDIT: removed content nested JSON property in responses
# EDIT: added ULID id generation

iex -S mix 
ContentApi.Repo.all(ContentApi.Api.Content)

mix ecto.reset
```

## Resources

* [go-sqlite3 - SQLite Library](https://github.com/mattn/go-sqlite3)
* [pgx - Postgres Library](https://github.com/jackc/pgx)
* [WAL File Size Issue with SQLite](https://news.ycombinator.com/item?id=40688987)

* [Huma Web Framework with OpenAPI Support](https://github.com/danielgtaylor/huma)
* [Huma Logging Middleware with Request Context](https://github.com/danielgtaylor/huma/blob/v1.14.3/middleware/logger.go)

