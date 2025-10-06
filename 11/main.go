package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Print("Количество значений в слайсах(пример: 4 5): ")

	var num1, num2 int

	_, err := fmt.Scan(&num1, &num2)
	if err != nil {
		fmt.Println("num1 или num2 не является числом")
		return
	}

	slice1 := make([]int, num1)
	slice2 := make([]int, num2)

	fillSlice(slice1, num1)
	fillSlice(slice2, num2)

	fmt.Printf("Слайс1: %v\nСлайс2: %v\n", slice1, slice2)

	set1 := toSet(slice1)
	set2 := toSet(slice2)

	res := makeIntersection(set1, set2)

	printResult(res)

}

func printResult(res map[int]struct{}) {
	counter := 0
	fmt.Print("Result: {")
	for key := range res {
		fmt.Printf("%d", key)
		counter++
		if counter != len(res) { // Постановка запятых между числами до последнего значения не включительно
			fmt.Print(", ")
		}
	}
	fmt.Println("}")
}

func fillSlice(s []int, countValues int) {
	seed := rand.NewSource(time.Now().UnixNano()) // Создаем уникальное стартовое значение для генератора чисел
	generator := rand.New(seed)                   // Генератор на основе сида

	for i := 0; i < countValues; i++ {
		s[i] = generator.Intn(8)
	}
}

func makeIntersection(first, second map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})

	for key := range first {
		if _, value := second[key]; value { // Если значение по данному ключу есть в другом слайсе, значит ключ попадает в мапу-результат
			result[key] = struct{}{}
		}
	}

	return result
}

func toSet(my_slice []int) map[int]struct{} {

	set := make(map[int]struct{})

	for _, element := range my_slice {
		set[element] = struct{}{}
	}

	return set
}
