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

func upload_content(sector string, base string, valor string, accion string) {
	var Rasputin_Port string
	var hostS string
	Rasputin_Port = ":50001"
	hostS = "dist149"

	connS, err := grpc.Dial(hostS+Rasputin_Port, grpc.WithInsecure()) //crea la conexion sincrona con el NameNode

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	service := pb.NewPlanetaryServiceClient(connS)

	switch accion {
	case "1":
		res, err := service.Add(context.Background(),
			&pb.BaseMessage{
				Sector: sector,
				Base:   base,
				Valor:  valor,
			})

		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}
		if res.Valor == "OK" {
			fmt.Printf("CORRECTO\n")
		}
	case "2":
		res, err := service.Rename(context.Background(),
			&pb.RenameMessage{
				Sector:  sector,
				Base:    base,
				Newbase: valor,
			})

		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}
		if res.Valor == "OK" {
			fmt.Printf("CORRECTO\n")
		}
	case "3":
		res, err := service.Delete(context.Background(),
			&pb.BaseMessage{
				Sector: sector,
				Base:   base,
				Valor:  valor,
			})

		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}
		if res.Valor == "OK" {
			fmt.Printf("CORRECTO\n")
		}
	case "4":
		res, err := service.Update(context.Background(),
			&pb.BaseMessage{
				Sector: sector,
				Base:   base,
				Valor:  valor,
			})

		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}
		if res.Valor == "OK" {
			fmt.Printf("CORRECTO\n")
		}

	default:
		connS.Close()
		return
	}

	connS.Close()

}

func main() {

	var menu string = "Eliga la accion a realizar:\n\t1:ADD\n\t2:RENAME\n\t3:DELETE\n\t4:UPDATE\n\tEleccion: "
	var election string
	var data string

	for {
		fmt.Print(menu)
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			election = scanner.Text()
		}

		fmt.Println("\nIngrese Input: ")
		if scanner.Scan() {
			data = scanner.Text()
		}

		lista := strings.Split(data, " ")
		if len(lista) == 2 {
			upload_content(lista[0], lista[1], "0", election)
		} else {
			upload_content(lista[0], lista[1], lista[2], election)
		}
	}
}
