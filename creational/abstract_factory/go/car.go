package abstract_factory

import (
	"errors"
	"fmt"
)

type Car interface {
	NumDoors() int
}

const (
	LuxuryCarType = iota
	FamilyCarType
)

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}
