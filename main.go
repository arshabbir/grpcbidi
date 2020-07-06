package main

import (
	"sync"

	"github.com/arshabbir/grpcbidirectional/client"
	"github.com/arshabbir/grpcbidirectional/server"
)

func main() {

	var wg sync.WaitGroup

	go server.Server.Start(&wg)

	go client.Client.Start(&wg)

	wg.Wait()
}
