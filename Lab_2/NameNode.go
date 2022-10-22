package main

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_2/Proto"
)

var Grunt_port string
var Synth_port string
var Cremator_port string

var hostQ string

func upload_content(DataNode_Port string, hostS string) {
	connS, err := grpc.Dial(hostS+DataNode_Port, grpc.WithInsecure()) //crea la conexion sincrona con el DataNode

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	service := pb.NewMessageServiceClient(connS)

	res, err := service.Upload(context.Background(),
		&pb.Message{
			Body: "Equipo listo?",
		})

}

func main() {
	hostQ = "localhost"
}
