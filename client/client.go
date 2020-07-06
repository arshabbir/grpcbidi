package client

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/arshabbir/grpcbidirectional/protopb"

	"google.golang.org/grpc"
)

type client struct {
	wg *sync.WaitGroup
}

type ClientService interface {
	Start(wg *sync.WaitGroup)
}

var (
	Client ClientService = &client{}
)

func (s *client) Send(stream protopb.MaxServiceBidi_MaxBidiClient, wg *sync.WaitGroup) {

	log.Println("Client send go routine started....")
	defer wg.Done()
	var i int = 0
	for {

		rand := rand.Intn(500)
		req := protopb.Request{Num: int64(rand)}
		i++
		if i >= 50 {
			break
		}
		stream.Send(&req)
		time.Sleep(time.Microsecond)
	}

}

func (s *client) Recv(stream protopb.MaxServiceBidi_MaxBidiClient, wg *sync.WaitGroup) {

	log.Println("Client recv go routine started....")
	defer wg.Done()
	for {

		msg, err := stream.Recv()
		if err != nil {
			log.Println("Error reading....")
			break
		}

		log.Println(msg)
		time.Sleep(time.Microsecond)
	}

}

func (s *client) Start(wg *sync.WaitGroup) {

	s.wg = wg

	defer wg.Done()
	log.Println("Starting client.....")

	cc, err := grpc.Dial("localhost:7077", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Error creating the client.....")
		return
	}

	client := protopb.NewMaxServiceBidiClient(cc)

	stream, err := client.MaxBidi(context.Background())
	if err != nil {

		log.Fatal("Error creating the clinet")
		return
	}

	var wg1 sync.WaitGroup

	wg1.Add(2)

	go s.Send(stream, &wg1)

	go s.Recv(stream, &wg1)

	wg1.Wait()
	//wg.Done()

}
