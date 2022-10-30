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

type fetchserver struct {
	pb.UnimplementedFetchServiceServer
}

var uploadServer *grpc.Server
var closeServer *grpc.Server
var fetchServer *grpc.Server

func CustomFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func appendtoFile(tipo string, id int, data string) {
	file, err := os.OpenFile("DATA.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
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
	err := os.Remove("DATA.txt")
	CustomFatal(err)
	return &pb.AckMessage{Ack: "OK"}, nil
}

func (s *fetchserver) Fetch(ctx context.Context, msg *pb.RequestToDataNodeMessage) (*pb.ReplyToNameNodeMessage, error) {
	fmt.Printf("Descarga solicitada: " + msg.Tipo + " " + msg.Id + "\n")
	fp, err := os.Open("DATA.txt")
	if err != nil {
		fmt.Println("Mensaje enviado: 0")
		return &pb.ReplyToNameNodeMessage{Si: "0"}, nil
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), ":")
		if splitLine[0] == msg.Tipo && splitLine[1] == msg.Id {
			identificador, err := strconv.Atoi(splitLine[1])
			CustomFatal(err)
			fmt.Println("Mensaje enviado; Tipo: " + msg.Tipo + "; Id: " + strconv.Itoa(identificador) + "; Data: " + splitLine[2])
			return &pb.ReplyToNameNodeMessage{Si: "1", Mensaje: &pb.Message{Tipo: msg.Tipo, Id: int64(identificador), Data: splitLine[2]}}, nil
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mensaje enviado: 0")
	return &pb.ReplyToNameNodeMessage{Si: "0"}, nil
}

func startCloseService(closeServer *grpc.Server, closeLis net.Listener) {
	pb.RegisterCloseServiceServer(closeServer, &closeserver{})
	if err := closeServer.Serve(closeLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
}

func startFetchService(fetchServer *grpc.Server, fetchLis net.Listener) {
	pb.RegisterFetchServiceServer(fetchServer, &fetchserver{})
	if err := fetchServer.Serve(fetchLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
}

func main() {
	uploadLis, err := net.Listen("tcp", ":50000")
	CustomFatal(err)
	closeLis, err := net.Listen("tcp", ":49000")
	CustomFatal(err)
	fetchLis, err := net.Listen("tcp", ":49500")
	CustomFatal(err)

	uploadServer = grpc.NewServer()
	closeServer = grpc.NewServer()
	fetchServer = grpc.NewServer()

	go startCloseService(closeServer, closeLis)

	go startFetchService(fetchServer, fetchLis)

	pb.RegisterUploadServiceServer(uploadServer, &uploadserver{})
	if err := uploadServer.Serve(uploadLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
}
