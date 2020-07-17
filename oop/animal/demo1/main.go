package main

//Go 面向对象编程（译）https://juejin.im/post/5d065cad51882523be6a92f2

import "fmt"

type Animal struct {
	Name string
	mean bool
}

/*我们可以使用 Dog 引用直接调用结构体 Animal 的成员，而 Cat 必须通过成员 Basics 访问到 Animal 的成员*/
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

type AnimalSounder interface {
	MakeNoise()
}

func MakeSomeNoise(animalSounder AnimalSounder) {
	animalSounder.MakeNoise()
}

func main() {
	fmt.Println("TEST")
	var dog Dog
	dog.mean = true
	dog.BarkStrength = 1
	dog.MakeNoise()
}
