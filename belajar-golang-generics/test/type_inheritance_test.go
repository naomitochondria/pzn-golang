package test

import (
	"fmt"
	"testing"
)

type Animal interface {
	EatingFood()
}

type Bird interface {
	EatingFood()
	Chirping()
}

type Parrot struct {
	Food string
}

func (parrot *Parrot) EatingFood() {
	fmt.Println("The parrot is eating some " + parrot.Food)
}

func (parrot *Parrot) Chirping() {
	fmt.Println("The parrot is chirping")
}

type Fish interface {
	EatingFood()
	Swimming()
}

type Dolphin struct {
	Food string
}

func (dolphin *Dolphin) EatingFood() {
	fmt.Println("The dolphin is eating some " + dolphin.Food)
}

func (dolphin *Dolphin) Swimming() {
	fmt.Println("The dolphin is swimming")
}

func FeedAnimal[T Animal](animal T) {
	fmt.Println("Let's bring the animal some food!")
	animal.EatingFood()
}

func TestTypeInherticance(t *testing.T) {
	FeedAnimal[*Parrot](&Parrot{Food: "Apple"})
	FeedAnimal[*Dolphin](&Dolphin{Food: "Squid"})
}
