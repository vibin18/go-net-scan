package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
)

type NetDevices struct {
	IP   string
	Name string
}

type mapping struct {
	Mac  string `yaml:"mac"`
	Name string `yaml:"name"`
}

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

func addDevicesToNetworkMap(ip net.IP, mac net.HardwareAddr) {
	myipString := fmt.Sprint(ip)
	mymacString := fmt.Sprint(mac)
	lock.Lock()
	NetworkDeviceMap[mymacString] = myipString
	lock.Unlock()
}

// map the devices to names comparing the mapping file

func mapDevices() {
	for mac, ip := range NetworkDeviceMap {
		for _, item := range MappedList {
			if mac == item.Mac {
				fmt.Printf("Adding %v as %v\n", item.Mac, item.Name)
				FinalMap[mac] = NetDevices{
					ip,
					item.Name,
				}
				fmt.Printf("Added %v\n", FinalMap[mac])
			} else {
				if mac == "00:11:32:a4:7f:2f" {
					fmt.Printf("Adding 00:11:32:a4:7f:2f as MAC")
				}
				FinalMap[mac] = NetDevices{
					ip,
					mac,
				}
			}

		}
	}
}
