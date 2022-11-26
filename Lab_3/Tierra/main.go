package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"strings"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_3/Proto"
	"google.golang.org/grpc"
)

type planetaryserver struct {
	pb.UnimplementedPlanetaryServiceServer
}

var planetaryServer *grpc.Server

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
	appendToFile(msg.Sector, msg.Base, msg.Valor)
	return &pb.ReplyMessage{Valor: "OK"}, nil
}

func (s *planetaryserver) Rename(ctx context.Context, msg *pb.RenameMessage) (*pb.ReplyMessage, error) {
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

func main() {
	planetaryLis, err := net.Listen("tcp", ":49000")
	CustomFatal(err)

	planetaryServer = grpc.NewServer()

	pb.RegisterPlanetaryServiceServer(planetaryServer, &planetaryserver{})
	if err := planetaryServer.Serve(planetaryLis); err != nil {
		panic("El servidor no se pudo iniciar" + err.Error())
	}
}
