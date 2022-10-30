package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_2/Proto"
)

func upload_content(tipo_data string, id int, data string) {
	var NameNode_Port string
	var hostS string
	NameNode_Port = ":50001"
	hostS = "dist149"

	connS, err := grpc.Dial(hostS+NameNode_Port, grpc.WithInsecure()) //crea la conexion sincrona con el NameNode

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
	if res.Ack == "NO" {
		fmt.Printf("ID REPETIDO, ERROR AL INTENTAR SUBIR MENSAJE\n")
	}
	if res.Ack == "OK" {
		fmt.Printf("Data enviada a NameNode\n")
	}
	connS.Close()

}

func main() {
	var data_type string
	var data_info string
	for true {
		fmt.Printf("ELIGA TIPO DE INFORMACION:\n")
		fmt.Printf("	1: MILITAR\n")
		fmt.Printf("	2: FINANCIERO\n")
		fmt.Printf("	3: LOGISTICO\n")
		fmt.Printf("	4: SALIR\n")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			data_type = scanner.Text()
		}
		if data_type == "4" {
			os.Exit(1)
		}
		if data_type == "1" {
			data_type = "MILITAR"
		}
		if data_type == "2" {
			data_type = "FINANCIERO"
		}
		if data_type == "3" {
			data_type = "LOGISTICO"
		}

		fmt.Print("INGRESE DATA DEL MENSAJE: ")
		if scanner.Scan() {
			data_info = scanner.Text()
		}
		fmt.Print("INGRESE ID DEL MENSAJE: ")
		if scanner.Scan() {
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				os.Exit(47)
			}
			upload_content(data_type, id, data_info)
		}

		//upload_content(data_type, id, data_info)
	}

}
