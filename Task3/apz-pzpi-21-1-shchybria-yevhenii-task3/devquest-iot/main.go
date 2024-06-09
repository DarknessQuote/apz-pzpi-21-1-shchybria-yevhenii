package main

import (
	"devquest-iot/device"
	"log"
)

func main() {
	currentDevice := device.GetDevice()
	
 	for _, _ = range make([]int, 10) {
		value, err := currentDevice.GetDataFromSensors()
		if err != nil {
			log.Panicln(err)
		} else {
			log.Println(value)
		}
 	}
}