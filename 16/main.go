package main

import (
	"fmt"
	"slices"
)

func main() {
	values := []int{8, 1, 19, 2, -4, -11, 9, 7, 6}

	fmt.Println("Before sorting:", values)
	fmt.Println("After sorting:", quickSort(values))
}

func quickSort(values []int) []int {

	if len(values) == 1 { // Если длина 1, возвращаем сразу число (слайс с 1 элементом)
		return values
	}

	target := values[len(values)/2] // Центральный элемент опорный -- такой метод эффективнее алгоритмов с 1-ым опорным элементом, но хуже чем с "медианой из трёх"

	lessThanTargetValues := make([]int, 0)
	moreThanTargetValues := make([]int, 0)

	for _, value := range values {
		if value > target {
			moreThanTargetValues = append(moreThanTargetValues, value) // Большие опорного в список больших
		} else if value < target {
			lessThanTargetValues = append(lessThanTargetValues, value) // Меньше опорного в список меньших
		}
	}

	return slices.Concat(quickSort(lessThanTargetValues), []int{target}, quickSort(moreThanTargetValues)) // Объедиянем и рекурсивно обрабатываем половины
}
