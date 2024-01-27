package network

func New(t TransmissionTypeable, i InterfaceTypeable) *NetworkType {
	return &NetworkType{
		transmission: t,
		innerface:    i,
	}
}

type NetworkTypeable interface {
	Transmission() TransmissionTypeable
	Interface() InterfaceTypeable
	// Other interface typeables...
}

var _ NetworkTypeable = (*NetworkType)(nil)

type NetworkType struct {
	transmission TransmissionTypeable
	innerface    InterfaceTypeable
}

func (n *NetworkType) Transmission() TransmissionTypeable {
	return n.transmission
}

func (n *NetworkType) Interface() InterfaceTypeable {
	return n.innerface
}
