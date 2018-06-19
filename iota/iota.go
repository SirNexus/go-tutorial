package main

import "fmt"

// Initialize iota iterator
// Set iota scaler at any point in code
const (
	_    = iota     // value: 0, iota: 0
	one             // value: 1, iota: 1
	four = 2 * iota // value: 4, iota: 2 - set scaler to iota * 2
	six             // value: 6, iota: 3
	_               // value: 8, iota: 4
	ten             // value: 10, iota: 5
)

func main() {
	fmt.Println("one:", one)
	fmt.Println("Ten:", ten)

	// Example mass variable declaration	
	var (
		testInt = 2
		testString = "Hello"
	)
	fmt.Println(testString, testInt)
}
