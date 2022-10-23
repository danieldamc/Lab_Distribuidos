package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"math/rand"
	"time"

	"google.golang.org/grpc"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_2/Proto"
)

var Grunt_port string
var Synth_port string
var Cremator_port string

var uploadServer *grpc.Server

var RECIBIDO = "MENSAJE RECIBIDO"

type uploadserver struct {
	pb.UnimplementedUploadServiceServer
}

func upload_content(tipo_data string, id int, data string) {
	var DataNode_Port string
	var hostS string
	var eleccion = rand.Intn(3)
	var local_usado string
	if eleccion == 0 {
		DataNode_Port = ":50000"
		hostS = "localhost"
		local_usado = "Grunt"
	} else {
		if eleccion == 1 {
			DataNode_Port = ":50000"
			hostS = "localhost"
			local_usado = "Cremator"
		} else {
			DataNode_Port = ":50000"
			hostS = "localhost"
			local_usado = "Synth"
		}
	}

	connS, err := grpc.Dial(hostS+DataNode_Port, grpc.WithInsecure()) //crea la conexion sincrona con el DataNode

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	service := pb.NewUploadServiceClient(connS)

	res, err := service.Upload(context.Background(),
		&pb.Message{
			Tipo: tipo_data,
			Id:   int64(id),
			Data: data,
		})

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}
	if res.Ack == "OK" {
		connS.Close()
	}
	fmt.Printf("Mensaje enviado a " + local_usado + " exitosamente.\n")

}

func (s *uploadserver) Upload(ctx context.Context, msg *pb.Message) (*pb.AckMessage, error) {
	fmt.Printf(msg.Tipo + "\n")
	upload_content(msg.Tipo, int(msg.Id), msg.Data)
	return &pb.AckMessage{Ack: "OK"}, nil
}

func main() {
	rand.Seed(time.Now().Unix())
	go upload_content("MILITAR", 1, "LLEGADA DE SUMINISTROS A DEPOSITO CITADELA")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			print(sig)
			fmt.Printf("\nctrl+c: Iniciando Protocolo de Cierre...\n")
			ConnClose, err := grpc.Dial("localhost:49000", grpc.WithInsecure())
			if err != nil {
				panic("No se pudo conectar con el servidor" + err.Error())
			}
			ServiceClose := pb.NewCloseServiceClient(ConnClose)
			ServiceClose.Close(context.Background(), &pb.CloseMessage{Close: "CLOSE"})
			ConnClose.Close()
			os.Exit(0)
		}
	}()

	uploadLis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatal("Error al escuchar en el puerto 50001")
	}
	uploadServer = grpc.NewServer()

	pb.RegisterUploadServiceServer(uploadServer, &uploadserver{})
	if err := uploadServer.Serve(uploadLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
}
