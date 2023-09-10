package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) int { // equivalent to : func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) { // x and y will be returned
	x = sum * 4 / 9
	y = sum - x
	return // this is called naked return
}

var c, python, java bool // var declares list of variables with same type at last

var i, j int = 1, 2 // A var declaration can include initializers, one per variable.

// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

const Pi = 3.14 // this is a constant

func needInt(x int) int { return x*10 + 1 }

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	fmt.Println("My favorite number is:", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(1, 1))
	a, b := swap("hello", "world") // multiple returns
	fmt.Println(a, b)
	var i int
	println(i, c, python, java)              //output: 0 false false false
	var c, python, java = true, false, "no!" // If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	fmt.Println(i, j, c, python, java)
	// k := 3 //Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

	// Type Conversions
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	fmt.Println(needInt(Small))
}
