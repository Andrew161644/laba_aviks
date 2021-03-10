package main

import (
	"flag"
	. "github.com/Andrew161644/full_app/grpc_service/grpc"
	. "github.com/Andrew161644/full_app/grpc_service/grpc/database/providers"
	"google.golang.org/grpc"
	"log"
	"net"
)

// для запуска открываем в терминале и вводим
//go run main.go -host=localhost
func main() {
	s := grpc.NewServer()
	var host = flag.String("host", "db", "HTTP listen address")

	flag.Parse()
	log.Println("Use host: " + *host)

	var db, err = Connect(*host, 5432, "postgres", "postgres", "postgres")

	srv := GrpcServer{
		DataBase: db,
	}
	RegisterDataServiceServer(s, &srv)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
