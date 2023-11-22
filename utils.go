package main

import (
	"errors"
	"fmt"
	"net"
)

func getAllInterfaces() []net.Interface {
	allIfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	return allIfaces
}

func validateInterface(iface string) (net.Interface, error) {
	allIfaces := getAllInterfaces()
	for _, i := range allIfaces {
		fmt.Println("Checking interface : ", i)
		if iface == i.Name {
			return i, nil
		}
	}
	return net.Interface{}, errors.New("Interface not found ")
}
