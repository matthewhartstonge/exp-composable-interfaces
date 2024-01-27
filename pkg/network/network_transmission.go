package network

import "errors"

type TransmissionTypeable interface {
	Type() string
	Send(payload []byte) error
	Receive() ([]byte, error)
}

var _ TransmissionTypeable = (*TransmissionType)(nil)

type TransmissionData struct {
	Type string
}

type TransmissionType struct {
	TransmissionData
}

func (t TransmissionType) Type() string {
	return t.TransmissionData.Type
}

func (t TransmissionType) Send(payload []byte) error {
	return errors.New("not implemented")
}

func (t TransmissionType) Receive() ([]byte, error) {
	return nil, errors.New("not implemented")
}
