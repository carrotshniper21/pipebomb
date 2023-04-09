// main.go
package main

import (
	"pipebomb/server"
)

func main() {
	port := "8001"
	server.StartServer(port)
}
