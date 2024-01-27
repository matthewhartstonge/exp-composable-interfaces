package physical

import (
	"github.com/matthewhartstonge/composable-interfaces/pkg/network"
)

func New(name string, hwAddr string) *Physical {
	return &Physical{
		InterfaceType: network.InterfaceType{
			InterfaceData: network.InterfaceData{
				Index: 1,
				MTU:   1450,
				Name:  name,
			},
		},
		hwAddr: hwAddr,
	}
}

type Physical struct {
	network.InterfaceType

	hwAddr string
}

func (p Physical) HardwareAddr() string {
	if p.hwAddr == "" {
		return "0xDEADBEEF"
	}
	return p.hwAddr
}
