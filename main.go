// pipebomb/main.go
package main

import (
	"fmt"
	"pipebomb/cache"
	"pipebomb/server"
)

//	@title			Pipebomb API
//	@version		6.9
//	@description	Pipebomb API for searching and streaming movies
//	@termsOfService	http://ani-j.netlify.app/tos/
//	@contact.name	API Support
//	@contact.url	https://github.com/ani-social
//	@contact.email	support@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			127.0.0.1:8001
//	@BaseURL		127.0.0.1:8001
//	@BasePath		/api
//	@schemes		https http
//
// pipebomb/main.go
func main() {
	port := "8001"
	redisCache := cache.NewCache("localhost:6379", "", 0) // replace with your Redis server details
	defer func(redisCache *cache.RedisCache) {
		err := redisCache.Close()
		if err != nil {
			fmt.Println("failed to close redis cache connection")
		}
	}(redisCache)
	server.StartServer(port, redisCache)
}
