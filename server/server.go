package server

import (
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/arshabbir/grpcbidirectional/protopb"

	//C:\Users\shabbir.hussain.NETXCELL\go\src\github.com\arshabbir\grpcbidi\protopb
	"google.golang.org/grpc"
)

type server struct {
	wg   *sync.WaitGroup
	nums []int64
}

type ServerService interface {
	Start(wg *sync.WaitGroup)
	MaxBidi(protopb.MaxServiceBidi_MaxBidiServer) error
}

var (
	Server ServerService = &server{}
)

func (s *server) Recv(stream protopb.MaxServiceBidi_MaxBidiServer, wg *sync.WaitGroup) {
	log.Printf(" Recv Go routing Started ......")
	var i int = 0
	defer wg.Done()

	for {

		msg, err := stream.Recv()

		if err == io.EOF {
			log.Printf("Stream Ended ....")

			return
		}
		if err != nil {
			log.Fatal("Error reading the stream...")
			return

		}
		log.Println(msg)
		s.nums[i] = msg.GetNum()
		i++
		_, max := MinMax(s.nums)

		resp := &protopb.Response{Maxstream: max}

		stream.Send(resp)
		time.Sleep(time.Microsecond)

	}

}



func (s *server) MaxBidi(stream protopb.MaxServiceBidi_MaxBidiServer) error {

	var wg sync.WaitGroup

	//Fork receive & send routines

	//go s.Send(stream, &wg)

	wg.Add(1)
	go s.Recv(stream, &wg)

	wg.Wait()

	return nil
}

func (s *server) Start(wg *sync.WaitGroup) {

	s.wg = wg
	s.nums = make([]int64, 100)

	lis, err := net.Listen("tcp", ":7077")
	if err != nil {
		log.Fatal("Error listining")
		return
	}

	server := grpc.NewServer()

	protopb.RegisterMaxServiceBidiServer(server, Server)

	log.Println("Starting gRPC server......")

	defer s.wg.Done()
	if err := server.Serve(lis); err != nil {
		log.Fatal("Listen Error")
	}
	return

}

func MinMax(array []int64) (int64, int64) {
	var max int64 = array[0]
	var min int64 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
