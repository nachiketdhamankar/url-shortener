module nachiket.me/url-shortener

go 1.16

replace nachiket.me/url-shortener/api => ./api

replace nachiket.me/url-shortener/mongodb => ./src/repository/mongodb

replace nachiket.me/url-shortener => ./shortener

require (
	github.com/go-chi/chi v1.5.4
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/onsi/ginkgo v1.16.1 // indirect
	github.com/onsi/gomega v1.11.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/teris-io/shortid v0.0.0-20201117134242-e59966efd125
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	go.mongodb.org/mongo-driver v1.5.1
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/dealancer/validate.v2 v2.1.0
)
