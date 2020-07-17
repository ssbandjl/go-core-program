package main

import "fmt"

type Animal struct {
	Name string
	mean bool
}

type Cat struct {
	Basics       Animal
	MeowStrength int
}

type Dog struct {
	Animal
	BarkStrength int
}

func (dog *Dog) MakeNoise() {
	barkStrength := dog.BarkStrength

	if dog.mean == true {
		barkStrength = barkStrength * 5
	}

	for bark := 0; bark < barkStrength; bark++ {
		fmt.Printf("BARK ")
		// fmt.Println("BARK ")
	}

	fmt.Println("")
}

func (cat *Cat) MakeNoise() {
	meowStrength := cat.MeowStrength

	if cat.Basics.mean == true {
		meowStrength = meowStrength * 5
	}

	for meow := 0; meow < meowStrength; meow++ {
		fmt.Printf("MEOW ")
	}

	fmt.Println("")
}

func main() {
	fmt.Println("TEST")
	var dog Dog
	dog.mean = true
	dog.BarkStrength = 1
	dog.MakeNoise()
}
