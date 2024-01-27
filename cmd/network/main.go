package main

import (
	"fmt"
	"github.com/matthewhartstonge/composable-interfaces/pkg/interface/physical"
	"github.com/matthewhartstonge/composable-interfaces/pkg/network"
	"github.com/matthewhartstonge/composable-interfaces/pkg/transmission/bluetooth"
)

func main() {
	nic := network.New(
		bluetooth.New(bluetooth.BLEEnabled),
		physical.New("bnet0", "0xB10E70074"),
	)

	fmt.Println("inet:", nic.Interface().Name())
	fmt.Println("inet type:", nic.Transmission().Type())

	// Here, PHY hasn't implemented Send() and the base "not implemented type"
	// methods are called instead.
	fmt.Println("send error:", nic.Transmission().Send([]byte("Hello, World!")))
}
