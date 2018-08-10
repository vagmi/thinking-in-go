package main

import (
	"fmt"
	"reflect"
)

// START_OMIT
type Person struct {
	FullName string `json:"full_name" mapsto:"fname"`
	Age      int    `json:"age" mapsto:"age"`
}

func main() {
	dent := Person{FullName: "Arthur Dent", Age: 36}
	structType := reflect.TypeOf(dent)
	val := reflect.ValueOf(dent)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if mappedName, ok := field.Tag.Lookup("mapsto"); ok {
			fmt.Printf(
				"%s mapped to %s with value of %v\n",
				field.Name,
				mappedName,
				val.FieldByName(field.Name))
		}
	}
}

// END_OMIT
