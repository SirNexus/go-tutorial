package main

import (
	"fmt"
)

func main() {
	m := map[string]bool{
		"SirNexus": true,
		"anonymous": false,
		"friend": true,
	}

	fmt.Printf("SirNexus: %v\n", m["SirNexus"])
	fmt.Printf("anonymous: %v\n", m["anonymous"])

	// "An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map."
	fmt.Printf("fakeKey: %v\n", m["fakeKey"])

	// Or you can test for existence by using multiple assignment
	key, ok := m["fakeKey"]
	fmt.Printf("exists: %v, key: %v\n", ok, key)

	fmt.Printf("\nDelete Key\n")
	key, ok = m["SirNexus"]
	fmt.Printf("SirNexus - exists: %v, key: %v\n", ok, key)

	// Delete key	
	delete(m, "SirNexus")

	// Alternatively don't care about key with use of "_"
	_, ok = m["SirNexus"]
	fmt.Printf("SirNexus - exists: %v\n", ok)
}
