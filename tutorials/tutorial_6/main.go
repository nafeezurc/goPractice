// structs and interfaces

package main

import (
	"fmt"
)

// to create a struct, use this syntax
type gasEngine struct {
	// and declare fields within body
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

// creating different types
type electricEngine struct {
	mpkwh     uint8
	kwh       uint8
	ownerInfo owner
}

// can even have another struct as a field in a struct
type owner struct {
	name string
}

// syntax for creating a method:
// inside paranthesis before method name identifies struct it operates with
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

// creating method for electric engine struct
func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// canMakeIt uses milesLeft methods to operate
// to make that function more general, you can use an interface
// create the interface, and have the function field that has the same method signature/return type (in this case, uint8)
// now anything with a method with that return type can work for the canMakeIt method
type engine interface {
	milesLeft() uint8
}

// function using structs
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Can make it")
	} else {
		fmt.Println("Can't make it")
	}
}

func main() {
	// can now declare a struct variable
	// by default fields get their default type values
	var myEngine gasEngine
	fmt.Println(myEngine.mpg, myEngine.gallons)

	// one way to declare field values is explicit
	myEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Println(myEngine.mpg, myEngine.gallons)

	// can also omit field names and fields are assigned in order
	myEngine = gasEngine{24, 16, owner{"Alex"}}
	fmt.Println(myEngine.mpg, myEngine.gallons)

	// can also set individual field values as follows
	myEngine.mpg = 26
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)

	// cam also create anonymous structs as follows
	// main differece is that this is not reusable
	// if you want another struct like this you have to recreate it
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{20, 20}
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	// using custom method for engine struct
	fmt.Printf("Miles left: %v\n", myEngine.milesLeft())

	// using fuction with struct
	canMakeIt(myEngine, 255)

	// after creating interface, can now use electric engine on canMakeIt as well
	var myEV electricEngine = electricEngine{25, 15, owner{"John"}}
	canMakeIt(myEV, 255)

}
