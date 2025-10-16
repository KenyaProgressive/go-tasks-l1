package main

import (
	"fmt"
	"slices"
)

func main() {
	sortedValues := make([]int, 0)

	for i := 1; i < 6; i++ {
		sortedValues = append(sortedValues, i)
	}

	// Считаем начальные границы слайса

	first := 0
	last := len(sortedValues) - 1

	target := 3

	index, status := binarySearch(sortedValues, target, first, last)
	indexByEmbeddedFunction, statusByEmbeddedFunction := slices.BinarySearch(sortedValues, target) // Выполним сравнение со встроенным бинарным поиском

	if status != "Success" { // Если ошибка -- выводим
		fmt.Println("My algorithm >>>", status)
		fmt.Println("Embedded >>>", statusByEmbeddedFunction)
	} else {
		fmt.Println("My algorithm >>> Value was found and have index:", index) // Иначе индекс найденного элемента
		fmt.Println("Embedded >>> Value was found and have index:", indexByEmbeddedFunction)
	}
}

func binarySearch(sortedValues []int, target int, first int, last int) (int, string) {
	if len(sortedValues) == 0 { // Если элементов нет -- сразу выдаем -1
		return -1, "No values in slice"
	} else if first > last { // Если нижная граница стала больше верхней -- значит значение не найдено
		return -1, "Value not found"
	} else {

		mid := (first + last) / 2 // считаем индекс центрального элемента
		// P.S. данный способ счета прост, но может быть небезопасен при больших диапазонах чисел -- может произойти переполнение памяти.

		if target > sortedValues[mid] { // В верхнюю половину, если искомый больше центрального
			return binarySearch(sortedValues, target, mid+1, last)
		} else if target < sortedValues[mid] { // В нижнюю половину, если искомый меньше центрального
			return binarySearch(sortedValues, target, first, mid-1)
		} else if target == sortedValues[mid] { // Если центральный и есть искомый -- возвращаем
			return mid, "Success"
		}
	}

	return -1, "Value not found" // -1 , если элемент не найден
}
