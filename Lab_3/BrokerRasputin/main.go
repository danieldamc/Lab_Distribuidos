package main

import (
	"context"
	"log"
	"net"

	"math/rand"
	"time"

	"google.golang.org/grpc"

	pb "github.com/danieldamc/Lab_Distribuidos/Lab_3/Proto"
)

var Tierra_port string
var Titan_port string
var Marte_port string

var PlanetaryServer *grpc.Server
var GetServer *grpc.Server

var RECIBIDO = "MENSAJE RECIBIDO"

type getserver struct {
	pb.UnimplementedGetServiceServer
}

type planetaryserver struct {
	pb.UnimplementedPlanetaryServiceServer
}

func CustomFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*
func appendtoFile(id int, nombre_datanode string) {
	file, err := os.OpenFile("DATA.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString(strconv.Itoa(id) + ":" + nombre_datanode + "\n"); err != nil {
		log.Fatal(err)
	}
}


func upload_content(tipo_data string, id int, data string) {

	var DataNode_Port string
	var hostS string
	var eleccion = rand.Intn(3)
	var local_usado string
	if eleccion == 0 {
		DataNode_Port = ":50000"
		hostS = "dist150"
		local_usado = "Grunt"
	} else {
		if eleccion == 1 {
			DataNode_Port = ":50000"
			hostS = "dist152"
			local_usado = "Cremator"
		} else {
			DataNode_Port = ":50000"
			hostS = "dist151"
			local_usado = "Synth"
		}
	}

	connS, err := grpc.Dial(hostS+DataNode_Port, grpc.WithInsecure())
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
	fmt.Printf("Mensaje de tipo " + tipo_data + " enviado a " + local_usado + " exitosamente.\n")
	appendtoFile(id, local_usado)

}

func closeDataNode(ip string, port string) {
	ConnClose, err := grpc.Dial(ip+port, grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	ServiceClose := pb.NewCloseServiceClient(ConnClose)
	ServiceClose.Close(context.Background(), &pb.CloseMessage{Close: "1"})
	ConnClose.Close()
}*/

func (s *getserver) Get(ctx context.Context, msg *pb.QueryMessage) (*pb.ReplyMessage, error) {
	mapa_conecciones := map[string]string{"Tierra": "dist150:49500", "Titan": "dist151:49500", "Marte": "dist152:49500"}

	/*
		connS, err := grpc.Dial(mapa_conecciones["Tierra"], grpc.WithInsecure())

		if err != nil {
			panic("No se pudo conectar con el servidor" + err.Error())
		}
		service := pb.NewGetServiceClient(connS)

		res, err := service.Get(context.Background(),
			&pb.QueryMessage{
				Sector: msg.Sector,
				Base:   msg.Base,
			})

		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}*/

	return &pb.ReplyMessage{Valor: mapa_conecciones["Tierra"]}, nil
}

func (s *planetaryserver) Add(ctx context.Context, msg *pb.BaseMessage) (*pb.ReplyMessage, error) {
	mapa_conecciones := map[string]string{"Tierra": "dist150:49500", "Titan": "dist151:49500", "Marte": "dist152:49500"}

	connS, err := grpc.Dial(mapa_conecciones["Tierra"], grpc.WithInsecure())

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	service := pb.NewPlanetaryServiceClient(connS)

	res, err := service.Add(context.Background(),
		&pb.BaseMessage{
			Sector: msg.Sector,
			Base:   msg.Base,
			Valor:  msg.Valor,
		})

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	return &pb.ReplyMessage{Valor: res.Valor}, nil
}

func (s *planetaryserver) Update(ctx context.Context, msg *pb.BaseMessage) (*pb.ReplyMessage, error) {
	mapa_conecciones := map[string]string{"Tierra": "dist150:49500", "Titan": "dist151:49500", "Marte": "dist152:49500"}

	connS, err := grpc.Dial(mapa_conecciones["Tierra"], grpc.WithInsecure())

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	service := pb.NewPlanetaryServiceClient(connS)

	res, err := service.Update(context.Background(),
		&pb.BaseMessage{
			Sector: msg.Sector,
			Base:   msg.Base,
			Valor:  msg.Valor,
		})

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	return &pb.ReplyMessage{Valor: res.Valor}, nil
}

func (s *planetaryserver) Rename(ctx context.Context, msg *pb.RenameMessage) (*pb.ReplyMessage, error) {
	mapa_conecciones := map[string]string{"Tierra": "dist150:49500", "Titan": "dist151:49500", "Marte": "dist152:49500"}

	connS, err := grpc.Dial(mapa_conecciones["Tierra"], grpc.WithInsecure())

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	service := pb.NewPlanetaryServiceClient(connS)

	res, err := service.Rename(context.Background(),
		&pb.RenameMessage{
			Sector:  msg.Sector,
			Base:    msg.Base,
			Newbase: msg.Newbase,
		})

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	return &pb.ReplyMessage{Valor: res.Valor}, nil
}

func (s *planetaryserver) Delete(ctx context.Context, msg *pb.BaseMessage) (*pb.ReplyMessage, error) {
	mapa_conecciones := map[string]string{"Tierra": "dist150:49500", "Titan": "dist151:49500", "Marte": "dist152:49500"}

	connS, err := grpc.Dial(mapa_conecciones["Tierra"], grpc.WithInsecure())

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	service := pb.NewPlanetaryServiceClient(connS)

	res, err := service.Update(context.Background(),
		&pb.BaseMessage{
			Sector: msg.Sector,
			Base:   msg.Base,
			Valor:  msg.Valor,
		})

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	return &pb.ReplyMessage{Valor: res.Valor}, nil
}

func startPlanetaryService(planetaryServer *grpc.Server, planetaryLis net.Listener) {
	pb.RegisterPlanetaryServiceServer(planetaryServer, &planetaryserver{})
	if err := planetaryServer.Serve(planetaryLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	planetaryLis, err := net.Listen("tcp", ":50001")
	getLis, err2 := net.Listen("tcp", ":49002")

	PlanetaryServer = grpc.NewServer()
	GetServer = grpc.NewServer()

	if err != nil {
		log.Fatal("Error al escuchar en el puerto 50001")
	}
	if err2 != nil {
		log.Fatal("Error al escuchar en el puerto 49002")
	}

	go startPlanetaryService(PlanetaryServer, planetaryLis)

	pb.RegisterGetServiceServer(GetServer, &getserver{})
	if err := GetServer.Serve(getLis); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}

}
