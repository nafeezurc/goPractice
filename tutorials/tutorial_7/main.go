// pointers

package main

import (
	"fmt"
)

func main() {
	// pointers are references to memory
	// to create a pointer, precede the type with *
	// by default, they point to nil
	// to have it point to a memory location, use new keyword
	var p *int32 = new(int32)
	var i int32

	// to get the value at the memory location, use the * symbol again
	// this is known as dereferencing
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	// to change a value at a location, dereference the variable
	// this operates very similarly to C/C++
	*p = 10
	fmt.Printf("The new value p points to is: %v\n", *p)

	// can assign a pointers memory address to another variables memory address with &
	p = &i
	// now if you dereference p and change the value, the value of i also changes
	*p = 1
	fmt.Printf("p: %v, i: %v\n", *p, i)

	// when using a regular variable and setting it equal to another value or variable, it simply copies that value into the original variable's memory location
	var k int32 = 2
	i = k
	fmt.Printf("p: %v, i: %v\n", *p, i)

	// main exception of copy behavior of non pointer variables is slices
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	// changing a value in the copy also changed the original
	// this is because the underlying structure of slices are pointers in an array

	// when using functions, we see that the values point to different memory locations, this is pass by value, meaning a copy is passed into the function
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of the thing1 array is: %p\n", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("The result is: %v\n", result)

	// because the pass by value is very inefficient, we can use pointers to deal with only one instance of our input, no copy required
	// this is pass by reference, meaning we pass the memory location in, and the value changes for our original input variable
	fmt.Printf("The memory location of the thing1 array is: %p\n", &thing1)
	result = pointerSquare(&thing1)
	fmt.Printf("The result is: %v\n", result)
	fmt.Printf("The value of thing1 is: %v\n", thing1)
	// we also see thing1 changed due to pass by reference
}

// using functions with slices
func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("The memory location of the thing2 array is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

// again, pointers are useful since we don't create copies of data
func pointerSquare(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of the thing2 array is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
