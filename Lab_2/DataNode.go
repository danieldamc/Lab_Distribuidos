package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

var uploadServer *grpc.Server

func main() {
	uploadLis, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatal("Error al escuchar en el puerto 50000")
	}
	uploadServer = grpc.NewServer()

}
