package main

import (
	"devquest-iot/device"
	"devquest-iot/management"
	"devquest-iot/network"
	"log"
)

func main() {
	configInstance, err := management.GetConfig()
	if err != nil {
		log.Panicln(err)
		return
	}

	currentDevice := device.GetDevice(configInstance)
	
 	for range make([]int, 10) {
		value, err := currentDevice.GetDataFromSensors()
		if err != nil {
			log.Panicln(err)
			return
		} else {
			log.Println(value)
		}
 	}

	httpConnection := management.NewHttpConnection(configInstance)
	requestSender := network.NewRequestSender(httpConnection, configInstance)

	res, err := requestSender.RegisterOwner("grimerssy", "12345678")
	if err != nil {
		log.Panicln(err)
		return
	}

	log.Printf(res.Message)
}