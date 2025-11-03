package main

import "fmt"

func main() {
	p1 := NewPoint(4, 12)
	p2 := NewPoint(2, 58)

	fmt.Println("Distance between p1 and p2:", p1.Distance(p2))
}
