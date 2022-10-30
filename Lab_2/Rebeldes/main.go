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

var RECIBIDO = "MENSAJE RECIBIDO"

var downloadServer *grpc.Server

type downloadserver struct {
	pb.UnimplementedDownloadServiceServer
}

func retrieve_content(query string) {
	var NameNode_Port string
	var hostS string
	NameNode_Port = ":50002"
	hostS = "dist149"

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

func Close() {
	var CloseHost string = "dist149"
	var ClosePort string = ":49000"

	CloseConn, err := grpc.Dial(CloseHost+ClosePort, grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	CloseService := pb.NewCloseServiceClient(CloseConn)

	CloseService.Close(context.Background(), &pb.CloseMessage{Close: "1"})
	CloseConn.Close()
}

func main() {
	var menu string = "Eliga la accion a realizar:\n\t1:Obtener Informacion Militar\n\t2:Obtener Informacion Financiera\n\t3:Obtener Informacion Logistica\n\t4:Cerrar NameNode y Datanodes\n\tEleccion: "
	var election string

	for {
		fmt.Print(menu)
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			election = scanner.Text()
		}
		switch election {
		case "1":
			fmt.Println("\nObteniendo informacion militar... ")
			retrieve_content("MILITAR")
		case "2":
			fmt.Println("\nObteniendo informacion financiera...")
			retrieve_content("FINANCIERO")
		case "3":
			fmt.Println("\nObteniendo informacion logistica...")
			retrieve_content("LOGISTICO")
		case "4":
			fmt.Println("\nIniciando proceso de cierre del Namenode y Datanodes...")
			Close()
			fmt.Println("Se han cerrado el Namenode y los Datanodes")
			os.Exit(0)
		default:
			fmt.Println("\nEleccion incorrecta")
		}
	}
}
