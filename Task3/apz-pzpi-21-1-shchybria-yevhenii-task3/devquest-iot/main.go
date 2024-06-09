package main

import (
	"devquest-iot/device"
	"devquest-iot/management"
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
}