package main

import (
	"fmt"

	"google.golang.org/grpc"
)

var Grunt_port string
var Synth_port string
var Cremator_port string

var RECIBIDO = "MENSAJE RECIBIDO"

func retrieve_content(tipo_data string, id int, data string) {
	var NameNode_Port string
	var hostS string
	NameNode_Port = ":50002"
	hostS = "localhost"

	connS, err := grpc.Dial(hostS+NameNode_Port, grpc.WithInsecure()) //crea la conexion sincrona con el NameNode

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	fmt.Print(connS)

	/*
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
		}*/

}

func main() {
	go retrieve_content("MILITAR", 1, "LLEGADA DE SUMINISTROS A DEPOSITO CITADELA")
}
