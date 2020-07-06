package server

import (
	"log"
	"net"
	"sync"

	"github.com/arshabbir/grpcbidirectional/protopb"
	"google.golang.org/grpc"
)

type server struct {
	wg *sync.WaitGroup
}

type ServerService interface {
	Start(wg *sync.WaitGroup)
	MaxBidi(MaxServiceBidi_MaxBidiServer) error
}

var (
	Server ServerService = &server{}
)

func (s *server) Recv(stream MaxServiceBidi_MaxBidiServer) {

}

func (s *server) Send(stream MaxServiceBidi_MaxBidiServer) {

}

func (s *server) MaxBidi(stream MaxServiceBidi_MaxBidiServer) error {

	wg.Add(2)

	go s.RecvandSend(stream)

	return nil
}

func (s *server) Start(wg *sync.WaitGroup) {

	s.wg = wg

	lis, err := net.Listen("tcp", ":7077")
	if err != nil {
		log.Fatal("Error listining")
		return
	}

	server := grpc.NewServer()

	protopb.RegisterMaxServiceBidiServer(server, Server)

	log.Println("Starting gRPC server......")

	if err := server.Serve(lis); err != nil {
		log.Fatal("Listen Error")
	}
	return

}
