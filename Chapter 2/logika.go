package main

import "fmt"

func main() {
	var x = true
	var y = false
	// Operator AND (&&)
	resultAnd := x && y
	fmt.Printf("x && y = %t\n", resultAnd)
	// Operator OR (||)
	resultOr := x || y
	fmt.Printf("x || y = %t\n", resultOr)
	// Operator NOT (!)
	resultNotX := !x
	resultNotY := !y
	fmt.Printf("!x = %t\n", resultNotX)
	fmt.Printf("!y = %t\n", resultNotY)
}
