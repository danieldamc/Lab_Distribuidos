package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"google.golang.org/grpc"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_3/Proto"
)

func GetSoldados(sector string, base string) {
	var Rasputin_Port string
	var hostS string
	Rasputin_Port = ":49002"
	hostS = "localhost"

	connS, err := grpc.Dial(hostS+Rasputin_Port, grpc.WithInsecure()) //crea la conexion sincrona con el NameNode

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	service := pb.NewGetServiceClient(connS)

	res, err := service.Get(context.Background(),
		&pb.QueryMessage{
			Sector: sector,
			Base:   base,
		})

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}
	fmt.Printf(res.Valor + "\n")
	direccion_planeta := res.Valor
	connS.Close()

	//Ahora nos comunicamos con el planeta directamente
	connS, err = grpc.Dial(direccion_planeta, grpc.WithInsecure()) //crea la conexion sincrona con el NameNode

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	service = pb.NewGetServiceClient(connS)

	res, err = service.Get(context.Background(),
		&pb.QueryMessage{
			Sector: sector,
			Base:   base,
		})

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	fmt.Print("Numero de soldados: " + res.Valor + "\n")
	connS.Close()
	return
}

func main() {

	var menu string = "Eliga Sector y Base: "
	var data string

	for {
		fmt.Print(menu)
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			data = scanner.Text()
		}

		lista := strings.Split(data, " ")
		GetSoldados(lista[0], lista[1])
	}
}
