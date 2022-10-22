package main

import (
	"context"

	"google.golang.org/grpc"
	"math/rand"
	"time"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_2/Proto"
	"google.golang.org/grpc"
)

var Grunt_port string
var Synth_port string
var Cremator_port string


var RECIBIDO = "MENSAJE RECIBIDO"

func upload_content(tipo string, id int, data string) {
	var DataNode_Port string
	var hostS string
	eleccion = rand.Int(2)
	if eleccion == 0 {
		DataNode_Port = ":49000"
		hostS = "dist149"
	}else{
		if eleccion == 1 {
			DataNode_Port = ":49001"
			hostS = "dist150"
		}else{
			DataNode_Port = ":49002"
			hostS = "dist151"
		}
	}

	connS, err := grpc.Dial(hostS+DataNode_Port, grpc.WithInsecure()) //crea la conexion sincrona con el DataNode

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	service := pb.NewUploadServiceClient(connS)

	res, err := service.Upload(context.Background(),
		&pb.Message{
			type: tipo,
			ID: id,
			data: data,
		})

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	if res == RECIBIDO{
		connS.Close()
	}
}

func main() {
	rand.Seed(time.Now().Unix())
}
