package main

import (
	"errors"
	"fmt"
)

func main() {
	// calling functions, basically same as Java
	// parameter type is enforced so have to pass in proper types
	// i.e can't input int because function needs string
	var printValue string = "Hello World"
	printMe(printValue)

	// return types, with a multiple return example
	var numerator int = 11
	var denominator int = 0
	// types shown will be inferred based on return types
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v\n", result)
	} else {
		// Printf example uses %v for variables
		fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
	}

	// switch statement version of the above
	// break is implied
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v\n", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
	}

	// conditional switch statements:
	switch remainder {
	case -1:
		fmt.Printf(err.Error())
	case 0:
		fmt.Printf("The division was exact\n")
	case 1, 2:
		fmt.Printf("The divison was close\n")
	default:
		fmt.Printf("The division was not close\n")
	}

}

// declaring functions as shown:
func printMe(printValue string) {
	fmt.Println(printValue)
}

// return type follows method header
// can also have multiple return items using (type, type) as shown
// also error handling here, so have to return an error as well
func intDivision(numerator int, denominator int) (int, int, error) {
	// error is another type in go used for error handling
	// by default has value set to nil
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero\n")
		return -1, -1, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
