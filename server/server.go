// pipebomb/server/server.go
package server

import (
	handies "github.com/gorilla/handlers"
	"log"
	"net/http"
	"pipebomb/cache"
	_ "pipebomb/docs"

	"github.com/gorilla/mux"
)

func Server(r *mux.Router, port string) {
	cors := handies.CORS(
		handies.AllowedOrigins([]string{"*"}),
		handies.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handies.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)
	log.Fatal(http.ListenAndServe(":"+port, cors(r)))
}

func StartServer(port string, redisCache *cache.RedisCache) {
	r := InitRouter(redisCache)
	Server(r, port)
}
