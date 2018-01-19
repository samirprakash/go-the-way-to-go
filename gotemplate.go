// Define the template for a simple go program

// Define the package as the very first line
package main

// Add the import statements after the package declaration
import (
	"fmt"
)

// Define any constants, variables or types
const c = "C"

var a int

type customType struct {
	// Define custom type
}

// Define the init function. It is not required for every package
func init() {
	// Package initialization
}

// Define main
func main() {
	var a int
	func1()
	fmt.Println(a)
}

// Define the functions as they have been called in main()
func (t customType) method1() {
	// Function body
}

func func1() {
	// Function body
}
