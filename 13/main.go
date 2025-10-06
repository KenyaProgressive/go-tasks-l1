package main

import "fmt"

func main() {
	a := 16
	b := 4

	swapWithoutExtraVar(a, b)

}

func swapWithoutExtraVar(a, b int) {
	fmt.Printf("Before: a=%d, b=%d\n", a, b)

	// На примере 4 и 16

	a += b             // a = 4 + 16 = 20
	b = (b - a) * (-1) // b = (16 - 20) * (-1) = 4
	a -= b             // a = 20 - 4 = 16

	fmt.Printf("After: a=%d, b=%d\n", a, b)
}
