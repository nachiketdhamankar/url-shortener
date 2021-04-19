## URL Shortener

### How to run the shortener

#### Pre-Run
- Install redis following [these](https://redis.io/topics/quickstart) steps
- Start a redis server (by running `./redis-server` from `redis-stable/src/`)
- Set the following env values
    * `URL_DB`
    * `REDIS_URL`
    
- To use the default configuration,
    * Don't pass any arguments to `./redis-server`
    * `URL_DB`=redis
    * `REDIS_URL`=redis://localhost:6379
    
#### 1. Binary
1. `cd bin`
2. `./url_shortener`

#### 2. From the source code
1. `cd src`
2. `go run main.go`

##### You can find the postman collection (of 2 requests) [here](docs/URL Shortener.postman_collection.json)