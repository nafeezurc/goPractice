// declaring variables and constants

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// declaring variables
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	// have to specify number of bits for floats
	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	// can't perform operations on different types, but can cast
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	// go uses floor division
	// % sign for remainders
	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	// string types, string is lowercase
	// can use escape sequences e.g. \n within strings
	// can also concatenate strings like Java e.g. "Hello" + " " + "World"
	var myString string = "Hello World"
	fmt.Println(myString)

	// can get length of a string with len() function
	// note that this is the number of bytes, not number of characters
	// characters outside the vanilla ascii set takes more than a single byte
	// e.g. "A" gives length 1, gamma symbol character gives length 2
	// can import unicode/utf8 library and use RuneCountInString() to get character numbers
	fmt.Println(len("test"))
	fmt.Println(len("γ"))
	fmt.Println(utf8.RuneCountInString("γ"))

	// rune data type repesents characters
	// printing them gives a numerical ascii value for the character
	var myRune rune = 'a'
	fmt.Println(myRune) // will print 97

	// booleans are simple
	var myBoolean bool = false
	fmt.Println(myBoolean)

	// types can be inferred based on values
	var myVar = "text"
	fmt.Println(myVar)

	// can go a step further and do shorthand for variables using :=
	myVar2 := "text2"
	fmt.Println(myVar2)

	// can also create multiple variables in one line
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// try to explicitly declare types when things are not obvious

	// everything discussed earlier applies to constants as well
	// main difference is constants can't be changed once created
	// value must be specified explicitly
	const myConst string = "const value"
	fmt.Println(myConst)
}
