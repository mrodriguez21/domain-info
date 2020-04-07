package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

const (
	port string = ":6463"
)

func main() {
	connection := NewDBConnection()
	service := NewService(connection)
	defer connection.db.Close()

	log.Fatal(fasthttp.ListenAndServe(port, service.Router.Handler))
}
