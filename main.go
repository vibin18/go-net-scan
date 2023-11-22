package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"go-net-scan/internal"
	"log"
	"net"
	"os"
	"sync"
)

var (
	argparser *flags.Parser
	arg       opts.Params
)

func initArgparser() {
	argparser = flags.NewParser(&arg, flags.Default)
	_, err := argparser.Parse()

	// check if there is a parse error
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			fmt.Println()
			argparser.WriteHelp(os.Stdout)
			os.Exit(1)
		}
	}
}

func main() {

	initArgparser()

	allIfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	var myIface net.Interface
	for _, i := range allIfaces {
		fmt.Println("Checking interface : ", i)
		if arg.Iface == i.Name {
			myIface = i
		} else {
			panic("Interface not found!")
		}
	}

	var wg sync.WaitGroup

	wg.Add(1)
	// Start up a scan on each interface.
	go func() {
		defer wg.Done()
		if err := scan(&myIface); err != nil {
			log.Printf("interface %v: %v", myIface.Name, err)
		}
	}()

	wg.Wait()

}
