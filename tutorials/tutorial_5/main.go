// strings, runes, and bytes

package main

import (
	"fmt"
	"strings"
)

func main() {
	// using non ascii characters to understand strings deeper
	var myString = "résumé"
	fmt.Println(myString)

	// can access characters in string like an array
	var indexed = myString[0]
	fmt.Println(indexed)

	// can use %T to get type
	fmt.Printf("%v, %T\n", indexed, indexed)

	// iterate through string characters
	// notice how non ascii values skip two in indeces
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// lengths and indeces in strings are based on bytes, not characters
	// so the "résumé" string will have length 8
	fmt.Println(len(myString))

	// rather than dealing with the underlying byte array
	// declare a string as a rune array/slice
	var myString2 = []rune("résumé")

	// now you can get contiguous indeces
	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	// and proper length based on runes
	fmt.Println(len(myString2))

	// strings are immutable in go, so you can't modify them
	// but you can use concatenation
	var strSlice = []string{"g", "o", "l", "a", "n", "g"}
	var catStr string
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("%v\n", catStr)

	// this is wildly inefficient as it creates a new string every time
	// for these scenarios, use a string builder, just like Java
	// need to import "strings" package
	var strBuilder strings.Builder
	// instead of '+' operator, we use the WriteString method
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	// then use String() method to convert to string
	var catStr2 = strBuilder.String()
	fmt.Printf("%v\n", catStr2)

}
