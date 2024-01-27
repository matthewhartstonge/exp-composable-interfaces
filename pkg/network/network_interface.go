package network

type InterfaceTypeable interface {
	Index() int
	MTU() int
	Name() string
}

var _ InterfaceTypeable = (*InterfaceType)(nil)

type InterfaceData struct {
	Index int    // positive integer that starts at one, zero is never used
	MTU   int    // maximum transmission unit
	Name  string // e.g., "en0", "lo0", "eth0.100"
}

type InterfaceType struct {
	InterfaceData
}

func (i *InterfaceType) Index() int {
	return i.InterfaceData.Index
}

func (i *InterfaceType) MTU() int {
	return i.InterfaceData.MTU
}

func (i *InterfaceType) Name() string {
	return i.InterfaceData.Name
}
