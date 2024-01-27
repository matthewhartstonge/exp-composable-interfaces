package vlan

import (
	"github.com/matthewhartstonge/composable-interfaces/pkg/network"
)

func New(name string, protocol Protocol) *VLan {
	return &VLan{
		InterfaceType: network.InterfaceType{
			InterfaceData: network.InterfaceData{
				Index: 100,
				MTU:   1500,
				Name:  name,
			},
		},
		Protocol: protocol,
	}
}

// Protocol is the type of tunnelling protocol.
type Protocol int

const (
	Unknown Protocol = iota
	IPSec
	GRE
)

type VLan struct {
	network.InterfaceType

	Protocol Protocol
}
