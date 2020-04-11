package main

import (
	"flag"
	"log"

	"github.com/AubSs/fasthttplogger"
	"github.com/valyala/fasthttp"
)

var (
	port = flag.String("port", ":8090", "Port to listen to")
	addr = flag.String("addr", "postgresql://root@localhost:26257/domain_info?sslmode=disable", "The address of the database")
)

func main() {
	flag.Parse()

	connection := NewDBConnection(*addr)
	service := NewService(connection)
	defer connection.db.Close()

	server := &fasthttp.Server{
		Handler: fasthttplogger.Tiny(service.Handler),
		Name:    "FastHttpLogger",
	}

	log.Fatal(server.ListenAndServe(*port))
}
