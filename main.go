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

	netDevice := net.Interface{
		Name: arg.Iface,
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
