package main

import "fmt"
// print demonstrates that ...int takes an arbitrary number of input integers and turns them into an array
func print(x ...int) {
	fmt.Println(x)
}


func main() {
	s := []int{1, 2, 3}

	print(1, 2, 3)
	// s must have ... appended so that it fits the type of int for the function it is being passed to
	print(s...)
}
