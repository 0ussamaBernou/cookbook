package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string { // like __str__ method in python
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur", 34}
	b := Person{"Ben", 32}
	fmt.Println(a, b)

}
