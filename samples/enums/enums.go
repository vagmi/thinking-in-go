package main

import "fmt"

// START1_OMIT
type CharacterType int

const (
	Human    CharacterType = iota // 0
	Droid                         // 1
	Starship                      // 2
)

type Character struct {
	Name string
	Type CharacterType
}

// END1_OMIT

// START2_OMIT
func main() {
	chars := []Character{
		{"Luke Skywalker", Human},
		{"R2D2", Droid},
		{"C3PO", Droid},
		{"Millenium Falcon", Starship},
	}
	fmt.Println(chars)
}

// END2_OMIT
