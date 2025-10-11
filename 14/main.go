package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan string)

	incorrectData := []int{1, 2, 3}

	printType(4)
	printType("golang")
	printType(true)
	printType(ch1)
	printType(ch2)
	printType((incorrectData))

}

func printType(data interface{}) {

	var t string

	switch data.(type) { // Использование type-switch (Альтернатива -- reflect.TypeOf)
	case int:
		t = "int"
	case string:
		t = "string"
	case bool:
		t = "bool"
	case chan any: // Канал, принимающий любые значения являются уникальным типом от остальнымх каналов (и они тоже каждый уникальный)
		t = "chan any"
	case chan int:
		t = "chan int"
	case chan string:
		t = "chan string"
	case chan bool:
		t = "chan bool"
	default:
		t = "is not from available -- int, bool, string, chan (any, int, string, bool)"
	}

	fmt.Printf("Recievied data have type --- %s\n", t)
}
