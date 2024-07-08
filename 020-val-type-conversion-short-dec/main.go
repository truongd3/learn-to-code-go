package main

import (
	"fmt"
)

var w = 99

func main() {
	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y) // 42, int
	fmt.Printf("%v of type %T \n", z, z) // 42, float64

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m) // float32 because declared

	/*
		// this does not work!
		z = m
		fmt.Printf("%v of type %T \n", z, z)
	*/

	// this does work
	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)
}

// %T --> type
