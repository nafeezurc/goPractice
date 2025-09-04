// Generics

package main

import (
	"fmt"
)

// structs can also make use of generics
type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	// rather than creating a bunch of different functions for different types, we can use generics to consolidate them into using one common function
	// for example, for summing an array of numeric values and having this function accept different numeric types:
	var intSlice = []int{1, 2, 3}
	// for the generic function, you can optionally specify the type in [] following the name of the function
	fmt.Println(sumSlice[int](intSlice))

	// the type being passed in is inferred, so we'll leave it out
	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice(float32Slice))

	// using our generic function accepting any type
	fmt.Println(isEmpty(intSlice))
	fmt.Println(isEmpty(intSlice))
	fmt.Println(isEmpty([]float64{}))
	fmt.Println(isEmpty([]string{}))

	// using our generic struct
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}
	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh:   57.5,
			mpkwh: 4.17,
		},
	}

	fmt.Println(gasCar)
	fmt.Println(electricCar)
}

// to specify what generic types are acceptable, place them in [] followed by the generic type (e.g. T) and separate each concrete type with the '|' symbol
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// we can use the 'any' type to let a function know any type is acceptable
// this won't work with all functions, only ones where operations are available to all types
// for this example, determining if a slice is empty by checking if it's length is 0
func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// for further learning, you can look into generics for loading jsons into their appropriate structs
