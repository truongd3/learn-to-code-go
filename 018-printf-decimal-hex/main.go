package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	fmt.Printf("%v \t %b \t %#X\n", a, a, a)
	fmt.Printf("%v \t %b \t %#X\n", b, b, b)
	fmt.Printf("%v \t %b \t %#X\n", c, c, c)
	fmt.Printf("%v \t %b \t %#X\n", d, d, d)
	fmt.Printf("%v \t %b \t %#X\n", e, e, e)
	fmt.Printf("%v \t %b \t %#X\n", f, f, f) // 5   101   0X5

	eva := -1
	fmt.Printf("-1 as binary, %b \n", eva)      // still -1
	fmt.Printf("-1 as hexadecimal, %x \n", eva) // still -1
}
