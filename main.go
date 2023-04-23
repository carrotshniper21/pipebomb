// pipebomb/main.go
package main

import (
	"pipebomb/server"
)


// @title Pipebomb API
// @version 6.9
// @description Pipebomb API for searching and streaming movies
// @termsOfService http://ani-j.netlify.app/tos/
// @contact.name API Support
// @contact.url https://github.com/ani-social
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host anij.bytecats.codes
// @BaseURL anij.bytecats.codes 
// @BasePath /pipebomb/api
// @schemes https http
func main() {
	port := "8001"
	server.StartServer(port)
}
