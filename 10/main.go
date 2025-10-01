package main

import (
	"fmt"
	"math"
)

func main() {
	temperatureSlice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := sortingTemperatures(temperatureSlice)

	for key, val := range result { // Красивый вывод
		fmt.Printf("%d: %.2f ", key, val)
		if key <= len(temperatureSlice) {
			fmt.Print(" ")
		}
	}
	fmt.Println()

}

func sortingTemperatures(temperatureSlice []float64) map[int][]float64 {
	resultMap := make(map[int][]float64)
	for _, value := range temperatureSlice {
		key := 0
		if value < 0 { // Если меньше 0, округляем в сторону отрицательную
			key = int(math.Floor(value/10)) * 10
		} else {
			key = int(value/10) * 10
		}

		resultMap[key] = append(resultMap[key], value) // Добавляем ключ и значение к ниму (если ключ существует, число значений в слайсе ключа увеличится)
	}

	return resultMap
}
