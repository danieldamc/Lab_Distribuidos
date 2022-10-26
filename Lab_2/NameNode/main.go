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
var downloadServer *grpc.Server

var RECIBIDO = "MENSAJE RECIBIDO"

type uploadserver struct {
	pb.UnimplementedUploadServiceServer
}

type downloadserver struct {
	pb.UnimplementedDownloadServiceServer
}

func upload_content(tipo_data string, id int, data string) {
	var DataNode_Port string
	var hostS string
	var eleccion = rand.Intn(3)
	var local_usado string
	if eleccion == 0 {
		DataNode_Port = ":50000"
		hostS = "dist150"
		local_usado = "Grunt"
	} else {
		if eleccion == 1 {
			DataNode_Port = ":50000"
			hostS = "dist151"
			local_usado = "Cremator"
		} else {
			DataNode_Port = ":50000"
			hostS = "dist152"
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

func (s *downloadserver) Download(ctx context.Context, msg *pb.RequestMessage) (*pb.ReplyMessage, error) {
	fmt.Printf("Descarga solicitada: " + msg.Tipo + "\n")
	var conecciones [3]string
	var n_mensajes int
	var mensajes_totales []*pb.Message
	conecciones[0] = "dist150:49500"
	conecciones[1] = "dist151:49500"
	conecciones[2] = "dist152:49500"
	for i := 0; i < 3; i++ {
		connS, err := grpc.Dial(conecciones[i], grpc.WithInsecure()) //crea la conexion sincrona con el DataNode

		if err != nil {
			panic("No se pudo conectar con el servidor" + err.Error())
		}

		service := pb.NewDownloadServiceClient(connS)

		res, err := service.Download(context.Background(),
			&pb.RequestMessage{
				Tipo: msg.Tipo,
			})

		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}
		fmt.Printf("Mensaje enviado exitosamente.\n")
		n_mensajes += int(res.Nmensajes)
		for j := 0; j < int(res.Nmensajes); j++ {
			mensajes_totales = append(mensajes_totales, res.Mensajes[j])
		}

	}

	return &pb.ReplyMessage{Nmensajes: int64(n_mensajes), Mensajes: mensajes_totales}, nil
}

func startDownloadService(downloadServer *grpc.Server, downloadLis net.Listener) {
	pb.RegisterDownloadServiceServer(downloadServer, &downloadserver{})
	if err := downloadServer.Serve(downloadLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
}

func main() {
	rand.Seed(time.Now().Unix())

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
	downloadLis, err2 := net.Listen("tcp", ":50002")

	uploadServer = grpc.NewServer()
	downloadServer = grpc.NewServer()

	if err != nil {
		log.Fatal("Error al escuchar en el puerto 50001")
	}

	if err2 != nil {
		log.Fatal("Error al escuchar en el puerto 50002")
	}

	go startDownloadService(downloadServer, downloadLis)

	pb.RegisterUploadServiceServer(uploadServer, &uploadserver{})
	if err := uploadServer.Serve(uploadLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}

}
