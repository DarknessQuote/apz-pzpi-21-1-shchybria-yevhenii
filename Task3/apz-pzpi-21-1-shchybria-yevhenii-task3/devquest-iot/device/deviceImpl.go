package device

import (
	"errors"
	"math/rand"
)

type Device struct {
	Type string
}

func GetDevice() IDevice {
	return &Device{Type: "pulse"}
}

func (d *Device) GetDataFromSensors() (int, error) {
	switch d.Type {
	case "pulse":
		return rand.Intn(100 - 60) + 60, nil
	default:
		return 0, errors.New("unsupported type of sensor")
	}
}