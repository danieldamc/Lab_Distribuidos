package main

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_2/Proto"
)

var RECIBIDO = "MENSAJE RECIBIDO"

var downloadServer *grpc.Server

type downloadserver struct {
	pb.UnimplementedDownloadServiceServer
}

func retrieve_content(query string) {
	var NameNode_Port string
	var hostS string
	NameNode_Port = ":50002"
	hostS = "localhost"

	connS, err := grpc.Dial(hostS+NameNode_Port, grpc.WithInsecure()) //crea la conexion sincrona con el NameNode

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	service := pb.NewDownloadServiceClient(connS)

	res, err := service.Download(context.Background(),
		&pb.RequestMessage{
			Tipo: query,
		})

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}
	if res.Nmensajes == -1 || res.Nmensajes == 0 { //ELEMENTO NO ENCONTRADO
		fmt.Printf("ELEMENTO NO ENCONTRADO\n")
	} else {
		fmt.Printf("Cantidad de mensaje descargados: " + strconv.Itoa(int(res.Nmensajes)) + "\n")
		for i := 0; i < int(res.Nmensajes); i++ {
			fmt.Printf(res.Mensajes[i].Data + "\n")
		}

	}
	connS.Close()

}

func main() {
	retrieve_content("MILITAR")
}
