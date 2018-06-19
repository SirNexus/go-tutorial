/*
The formatted print functions fmt.Fprint and friends take as a first argument any object that implements the io.Writer interface;
the variables os.Stdout and os.Stderr are familiar instances.

io.Writer interface:

type Writer interface {
        Write(p []byte) (n int, err error)
}

Must also define error type for custom error functionality.

Built-in error interface:

type error interface {
    Error() string
}

*/

package main

import (
	"fmt"
//	"io"
)

// Custom integer that combines the integer value with a string description of that value
type customInt struct {
	value int
	desc string
}

// Write is an implementation of the Write function from io.Writer to enable fmt.Fprintf to print using the custom struct
func (c *customInt) Write(p []byte) (n int, err error){
	c.desc = string(p)
	if c.desc == "" {
		return 0, c
	}
	return len(c.desc), nil;
}

// Error is the an implementation of the method within the error interface in order to make it compatible with the error type
func (c *customInt) Error() string {
	return fmt.Sprintf("That operation resulted in a value of: %v, desc: \"%v\", and resulted in an error.. Too bad :(", c.value, c.desc)
}

// print is a helper function to print all of the values of the customInt struct
func (c *customInt) Print() {
	fmt.Printf("CustomInt value: %v, description: %v\n", c.value, c.desc)
}


// Write to custom struct using implementation of various interfaces
func main() {
	/* This code block is an exploration of the inner workings of Fprintf when writing to an io.Writer interface type	
	Code left in for posterity

	p := make([]byte, 10)
	p = []byte("Test String")
	custom := customInt{}
	n, err := custom.Write(p)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("%v characters written\n", n)
	custom.Print()
	*/

	custom := customInt{3, "Initial String"}
	custom.Print()

	// Actual test of Fprintf functionality
	fmt.Fprintf(&custom, "This is a test using Fprintf!")
	custom.Print()
}

