package main

import "fmt"

type Car struct {
	year int
	make string
}

func (c *Car) String() string {
	return fmt.Sprintf("{make:%s, year:%d}", c.make, c.year)
}

func main() {
	myCar := Car{year: 2016, make: "Hyundai"}
	fmt.Println(myCar)
}
