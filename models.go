package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

var lock sync.Mutex

type NetDevices struct {
	IP   string
	MAC  string
	Name string
}

type mapping struct {
	Mac  string `yaml:"mac"`
	Name string `yaml:"name"`
}

var MappedList []mapping

func getConf(file string) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &MappedList)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

}

func addDevicesToDB(ip net.IP, mac net.HardwareAddr) {
	var wg sync.WaitGroup
	n := NetDevices{}
	db := []NetDevices{}
	myipString := fmt.Sprint(ip)
	mymacString := fmt.Sprint(mac)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, i := range MappedList {
			if mymacString == i.Mac {
				n.MAC = i.Mac
				n.Name = i.Name
				n.IP = myipString
			} else {
				n.MAC = i.Mac
				n.Name = i.Mac
				n.IP = myipString
			}
			fmt.Println("Adding ", n)
			db = append(db, n)
			lock.Lock()
			FinalList = &db
			lock.Unlock()
		}
	}()
	wg.Wait()
}
