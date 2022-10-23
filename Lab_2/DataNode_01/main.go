package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_2/Proto"
	"google.golang.org/grpc"
)

type uploadserver struct {
	pb.UnimplementedUploadServiceServer
}

var uploadServer *grpc.Server

func appendtoFile(tipo string, id int, data string) {
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString(tipo + ":" + strconv.Itoa(id) + ":" + data + "\n"); err != nil {
		log.Fatal(err)
	}
}

func (s *uploadserver) Upload(ctx context.Context, msg *pb.Message) (*pb.AckMessage, error) {
	fmt.Printf(msg.Tipo)
	appendtoFile(msg.Tipo, int(msg.Id), msg.Data)
	return &pb.AckMessage{Ack: "OK"}, nil
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
