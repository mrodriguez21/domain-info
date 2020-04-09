package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

const (
	port string = ":8090"
)

func main() {
	connection := NewDBConnection()
	service := NewService(connection)
	defer connection.db.Close()

	log.Fatal(fasthttp.ListenAndServe(port, service.Handler))
}
