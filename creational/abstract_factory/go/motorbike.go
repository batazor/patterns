package abstract_factory

import (
	"errors"
	"fmt"
)

type Motorbike interface {
	GetMotorbikeType() int
}

const (
	SportMotorbikeType  = iota
	CruiseMotorbikeType = iota
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}
