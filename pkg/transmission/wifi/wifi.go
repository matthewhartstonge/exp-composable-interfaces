package bluetooth

import (
	"github.com/matthewhartstonge/composable-interfaces/pkg/network"
)

// New returns a new bluetooth transmission type.
func New() *Wifi {
	return &Wifi{
		TransmissionType: network.TransmissionType{
			TransmissionData: network.TransmissionData{
				Type: "wifi",
			},
		},
	}
}

type Wifi struct {
	network.TransmissionType

	a  bool
	b  bool
	g  bool
	n  bool
	ax bool
}
