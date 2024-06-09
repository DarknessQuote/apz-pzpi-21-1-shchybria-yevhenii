package device

import (
	"devquest-iot/config"
	"errors"
	"math/rand"
	"sync"
)

type Device struct {
	Type string
}

var (
	once sync.Once
	deviceInstance *Device
)

func GetDevice(config *config.DeviceConfig) IDevice {
	once.Do(func() {
		deviceInstance = &Device{Type: config.DeviceSettings.Type}
	})

	return deviceInstance
}

func (d *Device) GetDataFromSensors() (int, error) {
	switch d.Type {
	case "pulse":
		return rand.Intn(100 - 60) + 60, nil
	default:
		return 0, errors.New("unsupported type of sensor")
	}
}