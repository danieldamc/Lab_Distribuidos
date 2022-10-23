package main

import (
	"context"

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
	DataNode_Port = ":50001"
	hostS = "localhost"

	connS, err := grpc.Dial(hostS+DataNode_Port, grpc.WithInsecure()) //crea la conexion sincrona con el NameNode

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
	go upload_content("MILITAR", 1, "LLEGADA DE SUMINISTROS A DEPOSITO CITADELA")
}
