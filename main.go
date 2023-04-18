// main.go
package main

import (
	"fmt"
	"pipebomb/logging"
	"pipebomb/server"

	"github.com/fatih/color"	
)

// @title Pipebomb API
// @version 1.0
// @description Pipebomb API for searching and streaming movies
// @BasePath /api
func main() {
	fmt.Println(color.GreenString(logging.HttpLogger()[0] + ":"), color.HiWhiteString(" Pipebomb running on http://127.0.0.1:8001 (Press CTRL+C to quit)"))
	port := "8001"
	server.StartServer(port)
}
