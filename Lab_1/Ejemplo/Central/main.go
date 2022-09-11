package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

var CASO_RESUELTO = "SI, ESTALLIDO CONTROLADO"
var qName string
var hostQ string
var hostS string

var ESCUADRONES_DISPONIBLES int

var port_lab1 string
var port_lab2 string
var port_lab3 string
var port_lab4 string

var m sync.Mutex

func resolver_estallido(port string, delivery amqp.Delivery) {
	m.Lock()
	ESCUADRONES_DISPONIBLES -= 1

	connS, err := grpc.Dial(hostS+port, grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	serviceCliente := pb.NewMessageServiceClient(connS)
	m.Unlock()
	for {
		//envia el mensaje al laboratorio
		time.Sleep(5 * time.Second) //espera de 5 segundos
		res, err := serviceCliente.Intercambio(context.Background(),
			&pb.Message{
				Body: "Equipo listo?",
			})

		//fmt.Printf(connS.GetState().String() + "\n")
		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}

		defer connS.Close() //defer cierra connS al final del for

		fmt.Println(string(delivery.Body) + " ha enviado: " + res.Body) //respuesta del laboratorio
		if res.Body == CASO_RESUELTO {
			fmt.Printf("Escuadron Retornando desde " + string(delivery.Body) + "...\n")
			connS.Close()
			m.Lock()
			ESCUADRONES_DISPONIBLES += 1
			m.Unlock()
			//fmt.Printf(connS.GetState().String() + "\n")
			break
		}

	}
}

func main() {
	ESCUADRONES_DISPONIBLES = 2
	var port string
	/*
		qName = "Emergencias"                                            //Nombre de la cola
		hostQ = "localhost"                                              //Host de RabbitMQ 172.17.0.1
		hostS = "localhost"                                              //Host de un Laboratorio
		connQ, err := amqp.Dial("amqp://guest:guest@" + hostQ + ":5672") //Conexion con RabbitMQ
	*/
	qName = "Emergencias" //Nombre de la cola
	hostQ = "localhost"   //Host de RabbitMQ 172.17.0.1
	//hostS = "localhost"                                              //Host de un Laboratorio
	connQ, err := amqp.Dial("amqp://guest:guest@" + hostQ + ":5672") //Conexion con RabbitMQ

	if err != nil {
		log.Fatal(err)
	}
	defer connQ.Close()

	ch, err := connQ.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(qName, false, false, false, false, nil) //amqp.Table{"x-max-length": 2}) //Se crea la cola en RabbitMQ
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)

	port_lab1 = ":50051" //puerto de la conexion con el laboratorio Pripyat 149
	port_lab2 = ":50052" //puerto de la conexion con el laboratorio Kampala 150
	port_lab3 = ":50053" //puerto de la conexion con el laboratorio Pohang 151
	port_lab4 = ":50050" //puerto de la conexion con el laboratorio Renca 152

	for {
		fmt.Println("Esperando Emergencias")
		chDelivery, err := ch.Consume(qName, "", true, false, false, false, nil) //obtiene la cola de RabbitMQ
		if err != nil {
			log.Fatal(err)
		}
		for delivery := range chDelivery {

			for {
				m.Lock()
				if ESCUADRONES_DISPONIBLES < 1 {
					m.Unlock()
					fmt.Printf("\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")
					fmt.Printf("Ningun Escuadron disponible,\nDebemos esperar que retorne uno\n")
					fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n")
					time.Sleep(time.Second * 5)
				} else {
					m.Unlock()
					break
				}
			}

			m.Lock()
			if ESCUADRONES_DISPONIBLES > 2 { //arreglar xd
				ESCUADRONES_DISPONIBLES = 2
			}

			fmt.Printf("--------------------------------\n")
			fmt.Println("numero de escuadrones disponibles: " + strconv.Itoa(ESCUADRONES_DISPONIBLES))
			fmt.Println("Pedido de ayuda de " + string(delivery.Body) + ". Enviando escuadron...") //obtiene el primer mensaje de la cola
			//fmt.Println(q)
			m.Unlock()

			if string(delivery.Body) == "Laboratorio Pripyat" {
				//fmt.Printf("pripyat momento\n")
				hostS = "localhost"
				port = port_lab1
				go resolver_estallido(port, delivery)
			}
			if string(delivery.Body) == "Laboratorio Kampala" {
				//fmt.Printf("kampala momento\n")
				hostS = "dist150"
				port = port_lab2
				go resolver_estallido(port, delivery)
			}
			if string(delivery.Body) == "Laboratorio Pohang" {
				//fmt.Printf("pohang momento\n")
				hostS = "dist151"
				port = port_lab3
				go resolver_estallido(port, delivery)
			}
			if string(delivery.Body) == "Laboratorio Renca" {
				//fmt.Printf("renca momento\n")
				hostS = "dist152"
				port = port_lab4
				go resolver_estallido(port, delivery)
			}

			//fmt.Println(port)

			/*
				connS, err := grpc.Dial(hostS+port, grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio

				if err != nil {
					panic("No se pudo conectar con el servidor" + err.Error())
				}

				serviceCliente := pb.NewMessageServiceClient(connS)

				for {
					//envia el mensaje al laboratorio
					res, err := serviceCliente.Intercambio(context.Background(),
						&pb.Message{
							Body: "Equipo listo?",
						})

					//fmt.Printf(connS.GetState().String() + "\n")
					if err != nil {
						panic("No se puede crear el mensaje " + err.Error())
					}

					defer connS.Close() //defer cierra connS al final del for

					fmt.Println(string(delivery.Body) + " ha enviado: " + res.Body) //respuesta del laboratorio
					if res.Body == CASO_RESUELTO {
						fmt.Printf("Escuadron Retornando...\n")
						connS.Close()
						//fmt.Printf(connS.GetState().String() + "\n")
						break
					}
					time.Sleep(5 * time.Second) //espera de 5 segundos

				}*/
		}
	}

}
