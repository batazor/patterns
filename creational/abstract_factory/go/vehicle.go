package abstract_factory

import (
	"errors"
	"fmt"
)

type Vehicle interface {
	NumWheels() int
	NumSeats() int
}

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType       = iota
	MotorbikeFactoryType = iota
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n", f))
	}
}
