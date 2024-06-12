package main

import (
	"devquest-iot/app"
	"devquest-iot/device"
	"devquest-iot/management"
	"fmt"
	"log"
)

func main() {
	configInstance, err := management.GetConfig()
	if err != nil {
		log.Panicln(err)
		return
	}

	currentDevice := device.GetDevice(configInstance)
	
 	dataProcessor := app.NewDataProcessor(currentDevice)
	avgValue, _ := dataProcessor.GetAverageValueFromSensors()
	log.Println(fmt.Sprintf("Average value: %f", avgValue))
}