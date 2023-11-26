package main

import (
	"encoding/json"
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

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// map the devices to names comparing the mapping file

func mapDevices() {
	for mac, ip := range NetworkDeviceMap {
		fmt.Printf(">Taking {#%v} from NetworkDeviceMap\n", mac)
		for _, item := range MappedList {
			fmt.Printf(">  Checking match with  {#%v} {#%v}\n", mac, item.Mac)
			if mac == item.Mac {
				fmt.Printf(">   	Match found!  {#%v} is {#%v}\n", mac, item.Mac)
				FinalMap[mac] = &NetDevices{
					ip,
					item.Name,
				}
				fmt.Printf(">   	Added {#%v} to FinalMap\n", item.Mac)
				fmt.Printf(">   		FinalMap[mac] is : {#%v} in FinalMap\n", FinalMap[item.Mac])
				break
			}
			fmt.Printf(">   	Match NOT  found!  {#%v} is NOT {#%v}\n", mac, item.Mac)
			FinalMap[mac] = &NetDevices{
				ip,
				mac,
			}

		}
		err := PrettyPrint(FinalMap)
		if err != nil {
			println(err)
		}
	}
}

// https://go.dev/play/p/w9NVq-931nY
