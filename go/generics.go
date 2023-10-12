package main

import "fmt"

type gasEngine struct {
	litter float32
	kmpL   float32
}
type electricEngine struct {
	kwh   float32 // KiloWattHour
	mpKwh float32 // MilePerKwh
}

type car[T electricEngine | gasEngine] struct {
	carMaker string
	carModel string
	engine   T
}

func main() {
	var intSlice = []int{1, 2, 3, 4, 5}
	var float32Slice = []float32{.3, .5, .2, .1, .4}

	var gasCar = car[gasEngine]{
		carMaker: "Ford",
		carModel: "F150",
		engine: gasEngine{
			litter: 15,
			kmpL:   30,
		},
	}

	fmt.Println(sumSlice[int](intSlice))
	fmt.Println(sumSlice[float32](float32Slice))
	fmt.Printf("is %T empty ? %v\n", intSlice, isEmpty(intSlice))

	fmt.Printf("gasCar: %+v", gasCar)

}

/* Function Generics */
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T

	for _, v := range slice {
		sum += v
	}

	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
