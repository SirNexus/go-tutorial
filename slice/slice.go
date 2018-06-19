package main

import (
	"fmt"
	"unsafe"
)

// printSlice prints the address of the slice, as well as what the slice references (underlying array),
// the memory footprint, size, capacity, and contents.
func printSlice(slice []int) {
	fmt.Printf("&Addr: %p, Addr: %p, mem: %v, Size: %v, Cap: %v %v\n", &slice, slice, unsafe.Sizeof(slice), len(slice), cap(slice), slice)
}

// printArr prints the address, memory footprint, size, capacity, and contents of a given array of length 5.
// Arrays can have their address printed with &arr, but printing arr with %p does not
// print a pointer. This is because, unlike C, arrays are values rather than pointers pointing
// to the head of the array. This can be seen by looking at the memory footprint of arrays vs slices
// Arrays have a memory footprint of 8 bits * num elements, while slices have a static footprint of 24 bits
func printArr(arr [5]int) {
	fmt.Printf("&Addr: %p, mem: %v, Size: %v, Cap: %v %v\n", &arr, unsafe.Sizeof(arr), len(arr), cap(arr), arr)
}

func main() {

	// Small array has less memory footprint than bigger array
	fmt.Println("Look at memory footprint of array vs slices\n")
	smallArr := [2]int{1,2}
	fmt.Printf("&Addr: %p, mem: %v, Size: %v, Cap: %v %v\n", &smallArr, unsafe.Sizeof(smallArr), len(smallArr), cap(smallArr), smallArr)
	var arr = [5]int{1, 2, 3, 4, 5}
	printArr(arr)
	slice := arr[:]
	printSlice(arr[3:])
	printSlice(slice)
	printSlice(smallArr[:])

	fmt.Println("\nMore Slices: Look how address is changed/incremented as slices are modified\n")

	slice = make([]int, 2)
	printSlice(slice)
	// Different slice location but references same underlying array
	printSlice(slice[:])

	// Referenced memory address incremented by one element. Different slice location
	printSlice(slice[1:])
	
	// Append creates totally new referenced array and slice
	slice = append(slice, 1)
	printSlice(slice[1:])
}
