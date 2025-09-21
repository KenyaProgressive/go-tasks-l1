package main

import (
	"fmt"
	"os"
)

func main() {
	var num int64       // Целевое число
	var shift int       // Сдвиг (номер бита)
	var newBitValue int // Значение бита

	fmt.Print("Введите число: ")

	if _, err := fmt.Fscan(os.Stdin, &num); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Введите номер бита (0..63): ")

	if _, err := fmt.Fscan(os.Stdin, &shift); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Введите новое значение бита (0 или 1): ")

	if _, err := fmt.Fscan(os.Stdin, &newBitValue); err != nil {
		fmt.Println(err)
		return
	}

	newNum := setBitInNumber(num, shift, newBitValue)

	if newNum != -1 {
		fmt.Printf("Число %d изменилось и стало %d, т.к. бит под номером %d был устанлвен в %d\n", num, newNum, shift, newBitValue)
	}

}

func setBitInNumber(num int64, shift int, value int) int64 {
	if (shift < 0) || (shift > 63) { // Если сдвиг некорректен -- меньше 0-го или больше 63 бита
		fmt.Println("Shift doesn't correct (0 <= shift <= 63)")
		return -1
	}

	if value == 1 { // Если установка бита в 1
		return (num | (1 << shift))
	} else if value == 0 { // Если установка бита в 0
		return (num &^ (1 << shift))
	} else { // Некорректное новое значение бита
		fmt.Println("Бит должен быть установлен в 0 или 1.")
		return -1
	}

}
