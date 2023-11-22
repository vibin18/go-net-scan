package main

import (
	"log"
	"net"
	"sync"
)

func main() {

	netDevice := net.Interface{
		Name: "en0",
	}

	var wg sync.WaitGroup

	wg.Add(1)
	// Start up a scan on each interface.
	go func() {
		defer wg.Done()
		if err := scan(&netDevice); err != nil {
			log.Printf("interface %v: %v", netDevice.Name, err)
		}
	}()

	wg.Wait()

}
