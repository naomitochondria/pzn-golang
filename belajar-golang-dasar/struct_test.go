package main

import (
	"fmt"
	"strconv"
	"testing"
)

type Bird interface {
	fly() string
}

type Parrot struct {
	Name           string
	SpecialitySong string
	FlyingSpeed    int
}

func (parrot Parrot) sing() string {
	return parrot.Name + " is singing " + parrot.SpecialitySong + " beautifully!"
}

func (parrot Parrot) fly() string {
	return parrot.Name + " is flying at " + strconv.Itoa(parrot.FlyingSpeed) + " km/h"
}

type Eagle struct {
	Name        string
	Prey        string
	FlyingSpeed int
}

func (eagle Eagle) hunting() string {
	return eagle.Name + " is hunting down those poor " + eagle.Prey
}

func (eagle Eagle) fly() string {
	return eagle.Name + " is flying at " + strconv.Itoa(eagle.FlyingSpeed) + " km/h"
}

func raceBird(birds ...Bird) {
	for _, bird := range birds {
		fmt.Println(bird.fly())
	}
}

func TestStruct(t *testing.T) {
	parrotA := Parrot{
		Name:           "Parrot A",
		SpecialitySong: "No Sleep",
		FlyingSpeed:    15,
	}

	eagleB := Eagle{
		Name:        "Eagle B",
		Prey:        "Mouse",
		FlyingSpeed: 60,
	}

	raceBird(parrotA, eagleB)

	var ikosong interface{} = 9
	fmt.Println(ikosong)
}
