package main

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

func main() {
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}
	swimmer.Train()
	swimmer.Swim()
}
