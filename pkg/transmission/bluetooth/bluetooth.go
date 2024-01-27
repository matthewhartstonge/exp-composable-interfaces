package bluetooth

import (
	"github.com/matthewhartstonge/composable-interfaces/pkg/network"
)

// New returns a new bluetooth transmission type.
func New(lowEnergy BLE) *Bluetooth {
	return &Bluetooth{
		TransmissionType: network.TransmissionType{
			TransmissionData: network.TransmissionData{
				Type: "bluetooth",
			},
		},
		LowEnergy: lowEnergy,
	}
}

type BLE int

const (
	BLEUnknown BLE = iota
	BLEDisabled
	BLEEnabled
)

type Bluetooth struct {
	network.TransmissionType

	LowEnergy BLE
}
