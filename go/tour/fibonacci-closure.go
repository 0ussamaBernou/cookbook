package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibs := []int{}
	return func() int {
		length := len(fibs)
		var fib int
		if length == 0 {
			fib = 0
		} else if length == 1 {
			fib = 1
		} else {
			fib = fibs[length-1] + fibs[length-2]
		}
		fibs = append(fibs, fib)
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
