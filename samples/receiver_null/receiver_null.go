package main

import "fmt"

// START_OMIT
type Animal interface {
	MakeNoise() string
}

type Dog struct {
	name string
}

func (d *Dog) MakeNoise() string {
	return fmt.Sprintf("%s says woof", d.name)
}

func main() {
	var animal Animal
	var d *Dog
	animal = d
	fmt.Println(animal.MakeNoise())
}

// END_OMIT
