package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_2/Proto"
	"google.golang.org/grpc"
)

type uploadserver struct {
	pb.UnimplementedUploadServiceServer
}

type closeserver struct {
	pb.UnimplementedCloseServiceServer
}

type downloadserver struct {
	pb.UnimplementedDownloadServiceServer
}

var uploadServer *grpc.Server
var closeServer *grpc.Server
var downloadServer *grpc.Server

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
	//err := os.Remove("file.txt")
	//CustomFatal(err)
	return &pb.AckMessage{Ack: "OK"}, nil
}

func (s *downloadserver) Download(ctx context.Context, msg *pb.RequestMessage) (*pb.ReplyMessage, error) {
	fmt.Printf("Descarga solicitada: " + msg.Tipo + "\n")
	fp, err := os.Open("file.txt")
	if err != nil {
		return &pb.ReplyMessage{
			Nmensajes: int64(-1),
		}, nil
	}
	defer fp.Close()
	var n int = 0
	//var rep *pb.ReplyMessage
	rep := &pb.ReplyMessage{Nmensajes: 0}
	var identificador int
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), ":")
		if splitLine[0] == msg.Tipo {
			identificador, err = strconv.Atoi(splitLine[1])
			CustomFatal(err)
			rep.Mensajes = append(rep.Mensajes, &pb.Message{
				Tipo: splitLine[0],
				Id:   int64(identificador),
				Data: splitLine[2],
			})
			n++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	rep.Nmensajes = int64(n)

	return rep, nil

}

func startCloseService(closeServer *grpc.Server, closeLis net.Listener) {
	pb.RegisterCloseServiceServer(closeServer, &closeserver{})
	if err := closeServer.Serve(closeLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
}

func startDownloadService(downloadServer *grpc.Server, downloadLis net.Listener) {
	pb.RegisterDownloadServiceServer(downloadServer, &downloadserver{})
	if err := downloadServer.Serve(downloadLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
}

func main() {
	uploadLis, err := net.Listen("tcp", ":50000")
	CustomFatal(err)
	closeLis, err := net.Listen("tcp", ":49000")
	CustomFatal(err)
	downloadLis, err := net.Listen("tcp", ":49500")
	CustomFatal(err)

	uploadServer = grpc.NewServer()
	closeServer = grpc.NewServer()
	downloadServer = grpc.NewServer()

	go startCloseService(closeServer, closeLis)

	go startDownloadService(downloadServer, downloadLis)

	pb.RegisterUploadServiceServer(uploadServer, &uploadserver{})
	if err := uploadServer.Serve(uploadLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
}
