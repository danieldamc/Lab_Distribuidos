package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

var serv *grpc.Server
var msg_intercambio string
var CASO_RESUELTO = "SI, ESTALLIDO CONTROLADO"
var CASO_NEGATIVO = "NO, ESTALLIDO AUN ACTIVO"
var listener net.Listener

func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	fmt.Println("La central dice: " + msg.Body)
	if rand.Float32() <= 0.6 {
		//defer serv.Stop()
		fmt.Printf("Situacion Resuelta!\n")
		msg_intercambio = CASO_RESUELTO
		return &pb.Message{Body: msg_intercambio}, nil
	}
	msg_intercambio = CASO_NEGATIVO
	return &pb.Message{Body: msg_intercambio}, nil
}

func empezarServicio(serv *grpc.Server, listener net.Listener) {
	pb.RegisterMessageServiceServer(serv, &server{})
	if err := serv.Serve(listener); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())

	LabName := "Laboratorio Renca" //nombre del laboratorio
	qName := "Emergencias"         //nombre de la cola
	hostQ := "152"                 //ip del servidor de RabbitMQ 172.17.0.1

	connQ, err := amqp.Dial("amqp://test:test@" + hostQ + ":5672") //conexion con RabbitMQ

	if err != nil {
		log.Fatal(err)
	}

	defer connQ.Close()
	ch, err := connQ.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	for {
		for {
			fmt.Printf("-----------------------\n")
			fmt.Printf("Todo bien... por ahora\n")
			time.Sleep(5 * time.Second)
			if rand.Float32() <= 0.8 {
				msg_intercambio = ""
				fmt.Printf("Estallido detectado!\n")
				//Mensaje enviado a la cola de RabbitMQ (Llamado de emergencia)
				//returns := ch.NotifyReturn(make(chan amqp.Return, 1))

				err = ch.Publish("", qName, false, false,
					amqp.Publishing{
						Headers:     nil,
						ContentType: "text/plain",
						Body:        []byte(LabName), //Contenido del mensaje
					})

				/*
					for r := range returns {
						fmt.Println(r)
					}
				*/
				if err != nil {
					fmt.Printf("error")
					log.Fatal(err)
				}
				break
			}
		}
		//fmt.Println(LabName)

		listener, err := net.Listen("tcp", ":50054") //conexion sincrona
		if err != nil {
			panic("La conexion no se pudo crear" + err.Error())
		}

		serv = grpc.NewServer()
		//defer serv.Stop()
		//for {
		//pb.RegisterMessageServiceServer(serv, &server{})
		go empezarServicio(serv, listener)

		/*
			if err = serv.Serve(listener); err != nil {
				panic("El server no se pudo iniciar" + err.Error())
			} */

		for {
			if msg_intercambio == CASO_RESUELTO {
				//fmt.Printf("ENTRA\n")
				time.Sleep(time.Second * 1)
				serv.Stop()
				listener.Close()
				break
			}
		}

	}

}
