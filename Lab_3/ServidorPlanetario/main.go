package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_3/Proto"
	"google.golang.org/grpc"
)

type getserver struct {
	pb.UnimplementedGetServiceServer
}

type planetaryserver struct {
	pb.UnimplementedPlanetaryServiceServer
}

var planetaryServer *grpc.Server
var GetServer *grpc.Server

func CustomFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func appendToFile(sector string, base string, nsoldados string) {
	file, err := os.OpenFile(sector+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString(sector + " " + base + " " + nsoldados + "\n"); err != nil {
		log.Fatal(err)
	}
}

func (s *planetaryserver) Add(ctx context.Context, msg *pb.BaseMessage) (*pb.ReplyMessage, error) {
	fmt.Println("Add; " + "Sector: " + msg.Sector + " Base: " + msg.Base + " nSoldados: " + msg.Valor)
	appendToFile(msg.Sector, msg.Base, msg.Valor)
	return &pb.ReplyMessage{Valor: "OK"}, nil
}

func (s *planetaryserver) Rename(ctx context.Context, msg *pb.RenameMessage) (*pb.ReplyMessage, error) {
	fp, err := os.Open(msg.Sector + ".txt")
	CustomFatal(err)

	fmt.Println("Rename; " + "Sector: " + msg.Sector + " Base: " + msg.Base + " NewBaseName: " + msg.Newbase)

	var lines []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fp.Close()

	for i, line := range lines {
		words := strings.Split(line, " ")
		if words[1] == msg.Base {
			lines[i] = msg.Sector + " " + msg.Newbase + " " + words[2]
		}
	}

	fp, err = os.OpenFile(msg.Sector+".txt", os.O_WRONLY|os.O_TRUNC, 0644)
	CustomFatal(err)

	for _, line := range lines {
		fp.WriteString(line + "\n")
	}
	fp.Close()

	return &pb.ReplyMessage{Valor: "OK"}, nil
}

func (s *planetaryserver) Update(ctx context.Context, msg *pb.BaseMessage) (*pb.ReplyMessage, error) {
	fp, err := os.Open(msg.Sector + ".txt")
	CustomFatal(err)

	fmt.Println("Update; " + "Sector: " + msg.Sector + " Base: " + msg.Base)

	var lines []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fp.Close()

	for i, line := range lines {
		words := strings.Split(line, " ")
		if words[1] == msg.Base {
			lines[i] = msg.Sector + " " + msg.Base + " " + msg.Valor
		}
	}

	fp, err = os.OpenFile(msg.Sector+".txt", os.O_WRONLY|os.O_TRUNC, 0644)
	CustomFatal(err)

	for _, line := range lines {
		fp.WriteString(line + "\n")
	}
	fp.Close()

	return &pb.ReplyMessage{Valor: "OK"}, nil
}

func (s *planetaryserver) Delete(ctx context.Context, msg *pb.BaseMessage) (*pb.ReplyMessage, error) {
	fmt.Println("Delete; " + "Sector: " + msg.Sector + " Base: " + msg.Base + " nSoldados: " + msg.Valor)

	fp, err := os.Open(msg.Sector + ".txt")
	CustomFatal(err)

	var lines []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fp.Close()

	for i, line := range lines {
		words := strings.Split(line, " ")
		if words[1] == msg.Base {
			lines[i] = " "
		}
	}

	fp, err = os.OpenFile(msg.Sector+".txt", os.O_WRONLY|os.O_TRUNC, 0644)
	CustomFatal(err)

	for _, line := range lines {
		if line != " " {
			fp.WriteString(line + "\n")
		}
	}
	fp.Close()
	return &pb.ReplyMessage{Valor: "SI"}, nil
}

func (s *getserver) Get(ctx context.Context, msg *pb.QueryMessage) (*pb.ReplyMessage, error) {
	//TODO: buscar en los txt numero de soldados de un sector y base
	fp, err := os.Open(msg.Sector + ".txt")
	CustomFatal(err)

	fmt.Println("Get; " + "Sector: " + msg.Sector + " Base: " + msg.Base)

	var lines []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fp.Close()

	for _, line := range lines {
		words := strings.Split(line, " ")
		if words[1] == msg.Base && words[0] == msg.Sector {
			return &pb.ReplyMessage{Valor: words[2]}, nil
		}
	}
	return &pb.ReplyMessage{Valor: "No existe: 0"}, nil
}

func startGetService(getServer *grpc.Server, getLis net.Listener) {
	pb.RegisterGetServiceServer(getServer, &getserver{})
	if err := getServer.Serve(getLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
}

func main() {
	planetaryLis, err := net.Listen("tcp", ":49900")
	getLis, err := net.Listen("tcp", ":60500")
	CustomFatal(err)

	planetaryServer = grpc.NewServer()
	GetServer = grpc.NewServer()

	go startGetService(GetServer, getLis)
	pb.RegisterPlanetaryServiceServer(planetaryServer, &planetaryserver{})
	if err := planetaryServer.Serve(planetaryLis); err != nil {
		panic("El servidor no se pudo iniciar" + err.Error())
	}
}
