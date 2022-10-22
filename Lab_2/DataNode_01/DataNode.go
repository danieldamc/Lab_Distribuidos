package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_2/Proto"
	"google.golang.org/grpc"
)

type uploadserver struct {
	pb.UnimplementedUploadServiceServer
}

var uploadServer *grpc.Server
var uploadLis net.Listener

func (s *uploadserver) Upload(ctx context.Context, msg *pb.Message) (*pb.AckMessage, error) {
	fmt.Println("La central dice: " + msg.Data)
	return &pb.AckMessage{ack: "OK"}, nil
}

func main() {
	uploadLis, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatal("Error al escuchar en el puerto 50000")
	}
	uploadServer = grpc.NewServer()

	pb.RegisterUploadServiceServer(uploadServer, &uploadserver{})
	if err := uploadServer.Serve(uploadLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}

}
