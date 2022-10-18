package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	pb "github.com/Sistemas-Distribuidos-2022-2/Tarea1-Grupo38/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

type server2 struct {
	pb.UnimplementedFinServiceServer
}

var serv *grpc.Server
var serv2 *grpc.Server
var msg_intercambio string
var CASO_RESUELTO = "SI, ESTALLIDO CONTROLADO"
var CASO_NEGATIVO = "NO, ESTALLIDO AUN ACTIVO"
var CASO_CIERRE = "cerrar"
var CASO_CIERRE_RESPUESTA = "ok"
var listener net.Listener
var listener2 net.Listener

func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {

	fmt.Println("La central dice: " + msg.Body)

	//calcula si el estallido es controlado
	if rand.Float32() <= 0.6 {
		fmt.Printf("Estallido contenido, Escuadron Retornando...\n")
		msg_intercambio = CASO_RESUELTO
		return &pb.Message{Body: msg_intercambio}, nil
	}
	msg_intercambio = CASO_NEGATIVO
	fmt.Printf("Revisando estado Escuadron: NO LISTO\n")
	return &pb.Message{Body: msg_intercambio}, nil
}

func (s *server2) Fin(ctx context.Context, msg *pb.Message2) (*pb.Message2, error) {
	defer os.Exit(1)
	fmt.Println("central ctrl+c")
	return &pb.Message2{Body: "OK"}, nil
}

func empezarServicio(serv *grpc.Server, listener net.Listener) {
	pb.RegisterMessageServiceServer(serv, &server{})
	if err := serv.Serve(listener); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
	return
}

func empezarServicio2(serv2 *grpc.Server, listener2 net.Listener) {
	pb.RegisterFinServiceServer(serv2, &server2{})
	if err := serv2.Serve(listener2); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())

	LabName := "Laboratorio Pohang" //nombre del laboratorio
	qName := "Emergencias"          //nombre de la cola
	hostQ := "dist149"              //ip del servidor de RabbitMQ 172.17.0.1

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

	listener2, err := net.Listen("tcp", ":49002")
	if err != nil {
		panic("La conexion no se pudo crear" + err.Error())
	}
	serv2 = grpc.NewServer()
	go empezarServicio2(serv2, listener2)

	for {
		for {
			fmt.Printf("-----------------------\n")
			fmt.Printf("Todo bien... por ahora\n")
			time.Sleep(5 * time.Second)
			if rand.Float32() <= 0.8 {
				msg_intercambio = ""
				fmt.Printf("Estallido detectado! SOS enviado a la central\n")
				//Mensaje enviado a la cola de RabbitMQ (Llamado de emergencia)

				err = ch.Publish("", qName, false, false,
					amqp.Publishing{
						Headers:     nil,
						ContentType: "text/plain",
						Body:        []byte(LabName), //Contenido del mensaje
					})

				if err != nil {
					fmt.Printf("error")
					log.Fatal(err)
				}
				break
			}
		}

		listener, err := net.Listen("tcp", ":50053") //conexion sincrona
		if err != nil {
			panic("La conexion no se pudo crear" + err.Error())
		}

		serv = grpc.NewServer()
		go empezarServicio(serv, listener)

		for {
			if msg_intercambio == CASO_RESUELTO {
				time.Sleep(time.Second * 1 / 100)
				serv.Stop()
				listener.Close()
				break
			}
		}

	}

}
