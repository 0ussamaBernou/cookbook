package main

import (
	"fmt"
)

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
}
func main() {
	c := make(chan int)

	go process(c)

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("exiting process")
}
