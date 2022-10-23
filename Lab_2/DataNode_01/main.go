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
type closeserver struct {
	pb.UnimplementedCloseServiceServer
}

var uploadServer *grpc.Server
var closeServer *grpc.Server
var CloseLis net.Listener

func CustomFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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
	fmt.Printf(msg.Tipo + "\n")
	appendtoFile(msg.Tipo, int(msg.Id), msg.Data)
	return &pb.AckMessage{Ack: "OK"}, nil
}

func (s *closeserver) Close(ctx context.Context, msg *pb.CloseMessage) (*pb.AckMessage, error) {
	defer os.Exit(0)
	fmt.Println("El Namenode esta cerrando, cerrando Datanode...")
	err := os.Remove("file.txt")
	CustomFatal(err)
	return &pb.AckMessage{Ack: "OK"}, nil
}

func startCloseService(closeServer *grpc.Server, closeLis net.Listener) {
	pb.RegisterCloseServiceServer(closeServer, &closeserver{})
	if err := closeServer.Serve(closeLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
}

func main() {
	uploadLis, err := net.Listen("tcp", ":50000")
	CustomFatal(err)
	closeLis, err := net.Listen("tcp", ":49000")
	CustomFatal(err)

	uploadServer = grpc.NewServer()
	closeServer = grpc.NewServer()

	go startCloseService(closeServer, closeLis)

	pb.RegisterUploadServiceServer(uploadServer, &uploadserver{})
	if err := uploadServer.Serve(uploadLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}

}
