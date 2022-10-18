package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	pb "github.com/Sistemas-Distribuidos-2022-2/Tarea1-Grupo38/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

var CASO_RESUELTO = "SI, ESTALLIDO CONTROLADO"

var CASO_CIERRE = "cerrar"
var CASO_CIERRE_RESPUESTA = "ok"

var qName string
var hostQ string
var hostS string

var ip string

//var ESCUADRONES_DISPONIBLES int

type ESCUADRA struct {
	nombre     string
	disponible bool
}

var ESCUADRA_UNO = ESCUADRA{"ESCUADRON ALFA", true}
var ESCUADRA_DOS = ESCUADRA{"ESCUADRON BETA", true}

var port_lab1 string
var port_lab2 string
var port_lab3 string
var port_lab4 string

var solicitudes [4]int
var labs [4]string

var m sync.Mutex

func handle_fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func resolver_estallido(port string, delivery amqp.Delivery, hostS string, escuadra ESCUADRA) {
	m.Lock()

	connS, err := grpc.Dial(hostS+port, grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	serviceCliente := pb.NewMessageServiceClient(connS)
	m.Unlock()
	for {

		//envia el mensaje al laboratorio
		time.Sleep(5 * time.Second) //espera de 5 segundos
		m.Lock()
		res, err := serviceCliente.Intercambio(context.Background(),
			&pb.Message{
				Body: "Equipo listo?",
			})

		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}

		defer connS.Close() //defer cierra connS al final del for

		fmt.Println(escuadra.nombre + " en " + string(delivery.Body) + " ha enviado: " + res.Body) //respuesta del laboratorio
		if res.Body == CASO_RESUELTO {

			fmt.Printf(escuadra.nombre + " retornando a la Central...")
			fmt.Println("Cerrando Conexion con " + string(delivery.Body) + "...")
			connS.Close()
			//liberamos la disponibilidad del escuadron usado
			if escuadra.nombre == ESCUADRA_UNO.nombre {
				ESCUADRA_UNO.disponible = true
			} else {
				ESCUADRA_DOS.disponible = true
			}
			m.Unlock()
			break
		}
		m.Unlock()

	}
}

func main() {

	solicitudes[0] = 0 //solicitudes hechas por laboratorio Pripyat
	solicitudes[1] = 0 //solicitudes hechas por laboratorio Kampala
	solicitudes[2] = 0 //solicitudes hechas por laboratorio Pohang
	solicitudes[3] = 0 //solicitudes hechas por laboratorio Renca

	labs[0] = "laboratorio Pripyat"
	labs[1] = "laboratorio Kampala"
	labs[2] = "laboratorio Pohang"
	labs[3] = "laboratorio Renca"

	var port string
	/*
		qName = "Emergencias"                                            //Nombre de la cola
		hostQ = "localhost"                                              //Host de RabbitMQ 172.17.0.1
		hostS = "localhost"                                              //Host de un Laboratorio
		connQ, err := amqp.Dial("amqp://guest:guest@" + hostQ + ":5672") //Conexion con RabbitMQ
	*/
	qName = "Emergencias"                                            //Nombre de la cola
	hostQ = "localhost"                                              //Host de RabbitMQ 172.17.0.1
	connQ, err := amqp.Dial("amqp://guest:guest@" + hostQ + ":5672") //Conexion con RabbitMQ
	handle_fatal(err)
	defer connQ.Close()

	ch, err := connQ.Channel()
	handle_fatal(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(qName, false, false, false, false, nil) //Se crea la cola en RabbitMQ
	handle_fatal(err)

	ch.QueuePurge(qName, false)
	//ch.QueueDelete()

	fmt.Println(q)

	port_lab1 = ":50051" //puerto de la conexion con el laboratorio Pripyat 149
	port_lab2 = ":50052" //puerto de la conexion con el laboratorio Kampala 150
	port_lab3 = ":50053" //puerto de la conexion con el laboratorio Pohang 151
	port_lab4 = ":50050" //puerto de la conexion con el laboratorio Renca 152

	puertos := [4]string{":49000", ":49001", ":49002", ":49003"}
	ip := [4]string{"dist149", "dist150", "dist151", "dist152"}

	fmt.Println("Esperando Emergencias")
	chDelivery, err := ch.Consume(qName, "", true, false, false, false, nil) //obtiene la cola de RabbitMQ
	handle_fatal(err)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("\n")
			print(sig)
			fmt.Printf("ctrl+c: Iniciando Protocolo de Cierre...\n")
			for i := 0; i < 4; i++ {
				connS2, err := grpc.Dial(ip[i]+puertos[i], grpc.WithInsecure())
				if err != nil {
					panic("No se pudo conectar con el servidor" + err.Error())
				}
				serviceCliente2 := pb.NewFinServiceClient(connS2)
				serviceCliente2.Fin(context.Background(), &pb.Message2{Body: "OK"})
				fmt.Println("Se Termino la Ejecucion del laboratorio en " + ip[i])
				//time.Sleep(time.Second)
				connS2.Close()
			}
			os.Exit(1)
		}
	}()

	for delivery := range chDelivery {
		for {
			m.Lock()
			//revisa si hay algun escuadron esta disponible
			if !ESCUADRA_UNO.disponible && !ESCUADRA_DOS.disponible {
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

		fmt.Printf("--------------------------------\n")

		if string(delivery.Body) == "Laboratorio Pripyat" {
			solicitudes[0] += 1
			hostS = "localhost"
			port = port_lab1
		} else if string(delivery.Body) == "Laboratorio Kampala" {
			solicitudes[1] += 1
			hostS = "dist150"
			port = port_lab2
		} else if string(delivery.Body) == "Laboratorio Pohang" {
			solicitudes[2] += 1
			hostS = "dist151"
			port = port_lab3
		} else if string(delivery.Body) == "Laboratorio Renca" {
			solicitudes[3] += 1
			hostS = "dist152"
			port = port_lab4
		}

		//asigna escuadra e inicia conexion con grpc
		if ESCUADRA_UNO.disponible {
			fmt.Println("Pedido de ayuda de " + string(delivery.Body) + ". Enviando " + ESCUADRA_UNO.nombre + "...")
			ESCUADRA_UNO.disponible = false
			m.Unlock()
			go resolver_estallido(port, delivery, hostS, ESCUADRA_UNO)
		} else {
			fmt.Println("Pedido de ayuda de " + string(delivery.Body) + ". Enviando " + ESCUADRA_DOS.nombre + "...")
			ESCUADRA_DOS.disponible = false
			m.Unlock()
			go resolver_estallido(port, delivery, hostS, ESCUADRA_DOS)
		}

		//escribir en el txt
		f, err := os.Create("SOLICITUDES.txt")
		handle_fatal(err)
		for i := 0; i < 4; i++ {
			_, err2 := f.WriteString(labs[i] + ";" + strconv.Itoa(solicitudes[i]) + "\n")
			handle_fatal(err2)
		}
		f.Close()
	}
}
