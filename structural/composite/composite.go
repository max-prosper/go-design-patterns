package main

import "fmt"

// ***** Direct Composition ***** //
type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swimming")
}

// ***** Embeddings Composition (example 1) ***** //
type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

// ***** Embeddings Composition (example 2) ***** //
type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

// ***** Embeddings Composition (example 3) ***** //
type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

func main() {
	// Direct
	swimmerA := CompositeSwimmerA{
		MySwim: Swim,
	}

	swimmerA.MyAthlete.Train()
	swimmerA.MySwim()

	// Embedded 1
	fish := Shark{
		Swim: Swim,
	}

	fish.Eat()
	fish.Swim()

	// Embedded 2
	swimmerB := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmerB.Train()
	swimmerB.Swim()

	// Embedded 3
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
			Left:      nil,
		},
		Left: &Tree{4, nil, nil},
	}

	fmt.Println(root.Right.Right.LeafValue)
}
