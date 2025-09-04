// arrays, slices, maps, and loops

package main

import (
	"fmt"
)

func main() {
	// declaring an array
	var intArr [3]int32

	//changing value of an element
	intArr[1] = 123
	// 0 indexed
	fmt.Println(intArr[0])
	// can access ranges, here it accesses elements 1 and 2
	fmt.Println(intArr[1:3])

	//'&' symbol can be used to access memory locations
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// explicit array declaration
	intArr2 := [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	// slices are wrappers for arrays that allows for some dynamic properties
	// note that it's not the same as a vector or list
	// more of a "look-into" an underlying array
	// simply omit size declaration to create one
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	// can access capacity of array using 'cap' function
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	// can add values to a slice with append function
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice[3])

	// can append multiple values by using the spread operator as shown
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// another way to create a slice is using make function
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// using maps, operate with key-value pairs
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	// explicit map declaration
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])

	// map always returns a value whether or not a key exists
	// if it doesn't, it returns a default value
	// you have to handle this issue yourself, since it always returns something, you can run into problems
	// to fix this, use map's built in boolean return value that returns true if a key exists
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid name")
	}

	// to delete from a map, use delete function
	delete(myMap2, "Adam")

	// iterating requires range keyword in for loop
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}
	// note that maps may not always return its elements in the same order every time
	// can get different orders on different calls

	// for arrays and slices
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// go doesn't have while loops, but can simulate by doing:
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// or
	var j int = 0
	for {
		if j >= 10 {
			break
		}
		fmt.Println(j)
		j++
	}

	// or
	for k := 0; k < 10; k++ {
		fmt.Println(k)
	}

}
