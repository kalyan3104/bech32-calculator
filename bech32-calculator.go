package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Define the hexadecimal number as a string
	hexadecimalNumber := "000000000000000000010000000000000000000000000000000000000001ffff"

	// Create a new big.Int to hold the decimal value
	decimalValue := new(big.Int)

	// Set the decimal value by parsing the hexadecimal string
	decimalValue, success := decimalValue.SetString(hexadecimalNumber, 16)

	if !success {
		fmt.Println("Failed to parse the hexadecimal number.")
		return
	}

	// Print the result
	fmt.Println("Decimal Value:", decimalValue)
}
