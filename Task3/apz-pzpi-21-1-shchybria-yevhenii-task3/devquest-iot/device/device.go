package device

type IDevice interface {
	GetDataFromSensors() (int, error)
}