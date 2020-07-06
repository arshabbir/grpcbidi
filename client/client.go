package client

import (
	"log"
	"sync"
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

func (s *client) Send() {

}

func (s *client) Recv() {

}

func (s *client) Start(wg *sync.WaitGroup) {
	log.Println("Starting client.....")

	s.wg = wg
}
