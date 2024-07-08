package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// this would not work because e is not used
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// this does work
	var g int
	fmt.Println(g) // default 0
	g = 43
	fmt.Println(g)

	// DECLARE a variable to hold a VALUE of a certain TYPE then ASSIGN a VALUE of that TYPE to the variable
	var h int = 44 // initializing a variable
	fmt.Println(h)
}
