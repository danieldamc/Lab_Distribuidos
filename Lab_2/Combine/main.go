package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_2/Proto"
)

func upload_content(tipo_data string, id int, data string) {
	fmt.Printf("ENVIANDO...\n")
	fmt.Printf(tipo_data)
	fmt.Printf("\n")
	fmt.Print(id)
	fmt.Printf("\n")
	fmt.Printf(data)
	fmt.Printf("\n")
	var NameNode_Port string
	var hostS string
	NameNode_Port = ":50001"
	hostS = "localhost"

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
	if res.Ack == "OK" {
		connS.Close()
	}
	fmt.Printf("Data enviada a NameNode\n")

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

		fmt.Printf("INGRESE DATA DEL MENSAJE\n")
		if scanner.Scan() {
			data_info = scanner.Text()
		}
		upload_content(data_type, 1, data_info)
	}

}
