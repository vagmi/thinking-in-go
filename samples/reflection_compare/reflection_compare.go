package main

import (
	"fmt"
	"reflect"
)

// START_OMIT
type Status struct {
	Status string
	Code   int
}

type StructA struct {
	Name   string
	Value  int
	Status Status
}

func main() {
	strA := &StructA{Value: 20, Name: "Go", Status: Status{Status: "OK", Code: 200}}
	strB := &StructA{Value: 20, Name: "Go", Status: Status{Status: "OK", Code: 200}}
	if reflect.DeepEqual(strA, strB) {
		fmt.Println("Yay!")
	}
}

// END_OMIT
