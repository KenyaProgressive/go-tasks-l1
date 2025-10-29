package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	startMenu()
}

func startMenu() {

	reader := bufio.NewReader(os.Stdin)
	a := new(big.Float) // Выделим память под числа a и b
	b := new(big.Float)

	fmt.Print("Введите число a: ")

	aStr, _ := reader.ReadString('\n') // Считывание значения a

	if _, ok1 := a.SetString(strings.TrimSpace(aStr)); !ok1 { // Если число некорректное -- выход из программы
		fmt.Println("Введённое число некорректно.")
		return
	}

	fmt.Print("Введите число b: ")

	bStr, _ := reader.ReadString('\n')

	if _, ok2 := b.SetString(strings.TrimSpace(bStr)); !ok2 {
		fmt.Println("Введённое число некорректно.")
		return
	}

	fmt.Println("Выберите номер операции: ")

	operation := showOperationsAndGetOneOfThem()

	if operation == '!' {
		fmt.Println("Недопустимая аримфетическая операция")
		return
	}

	calcBigNumbers(a, b, operation)

}

func calcBigNumbers(a *big.Float, b *big.Float, operation rune) {

	result := big.NewFloat(0.0)

	switch operation {
	case '+':
		fmt.Println("Результат сложения ---", result.Add(a, b))
	case '-':
		fmt.Println("Результат вычитания ---", result.Sub(a, b))
	case '*':
		fmt.Println("Результат умножения ---", result.Mul(a, b))
	case '/':
		if b.Sign() == 0 { // Знак нуля будет 0
			fmt.Println("Деление на 0 запрещено!")
		} else {
			fmt.Println("Результат деления ---", result.Quo(a, b))
		}
	}
}

func showOperationsAndGetOneOfThem() rune {

	fmt.Println("1 -- Сложение")
	fmt.Println("2 -- Вычитание")
	fmt.Println("3 -- Умножение")
	fmt.Println("4 -- Деление")

	var operationNumber int

	fmt.Scan(&operationNumber)

	return opNumberToOpSymbol(operationNumber)
}

func opNumberToOpSymbol(operationNumber int) rune {
	switch operationNumber {
	case 1:
		return '+'
	case 2:
		return '-'
	case 3:
		return '*'
	case 4:
		return '/'
	default:
		return '!'
	}
}
