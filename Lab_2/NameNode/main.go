package main

import (
	"context"
	"fmt"
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

var RECIBIDO = "MENSAJE RECIBIDO"

func upload_content(tipo_data string, id int, data string) {
	var DataNode_Port string
	var hostS string
	var eleccion = rand.Intn(3)
	if eleccion == 0 {
		DataNode_Port = ":50000"
		hostS = "localhost"
	} else {
		if eleccion == 1 {
			DataNode_Port = ":50000"
			hostS = "localhost"
		} else {
			DataNode_Port = ":50000"
			hostS = "localhost"
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

}

func main() {
	rand.Seed(time.Now().Unix())
	go upload_content("MILITAR", 1, "LLEGADA DE SUMINISTROS A DEPOSITO CITADELA")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			print(sig)
			fmt.Printf("ctrl+c: Iniciando Protocolo de Cierre...\n")
			ConnClose, err := grpc.Dial("localhost:49000", grpc.WithInsecure())
			if err != nil {
				panic("No se pudo conectar con el servidor" + err.Error())
			}
			ServiceClose := pb.NewCloseServiceClient(ConnClose)
			res, err := ServiceClose.Close(context.Background(), &pb.CloseMessage{Close: "CLOSE"})
			if err != nil {
				panic("No se puede crear el mensaje " + err.Error())
			}
			if res.Ack == "OK" {
				fmt.Println("Se termino la ejecucion del Datanode")
			}
			ConnClose.Close()
		}
	}()
}
