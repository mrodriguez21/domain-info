package main

import (
	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	port = flag.String("port", ":8090", "Port to listen to")
)

func main() {
	flag.Parse()

	connection := NewDBConnection()
	service := NewService(connection)
	defer connection.db.Close()

	log.Fatal(fasthttp.ListenAndServe(*port, service.Handler))
}
