package main

import (
	"flag"
	"log"

	"github.com/AubSs/fasthttplogger"
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

	server := &fasthttp.Server{
		Handler: fasthttplogger.Tiny(service.Handler),
		Name:    "FastHttpLogger",
	}

	log.Fatal(server.ListenAndServe(*port))
}
