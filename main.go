// main.go
package main

import (
	"pipebomb/server"
)

// @title Pipebomb API
// @version 1.0
// @description Pipebomb API for searching and streaming movies
// @BasePath /api
func main() {
	port := "8001"
	server.StartServer(port)
}
