package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"time"
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

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // if can start with a short statement
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}
func helloworld() {
	defer fmt.Println("world!") // the execution of this line will be defered until the surrounding function returns

	fmt.Println("Hello")
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

var pows = []int{1, 2, 4, 8, 16, 32, 64, 128}

type Coord struct {
	Lat, Long float64
}

var m map[string]Coord

var m2 = map[string]Coord{
	"Google": {
		37.42202, -122.08408,
	},
}

// function as value
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// function closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

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
	sum := 1
	for sum < 100 {
		sum += sum
	}

	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	helloworld()

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(v1, p, v2, v3)
	primes := [6]int{2, 3, 5, 7, 11, 13} // array

	var s []int = primes[1:4] // slice
	s[0] = 17
	fmt.Println(s, primes)
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	for i, v := range pows {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// maps
	m = make(map[string]Coord)
	m["Bell Labs"] = Coord{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	//functions as values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
