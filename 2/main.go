package main

import (
	"fmt"
	"sync"
)

func main() {

	massiv := []int{2, 4, 6, 8, 10}

	wgSquares := sync.WaitGroup{} // Создадим wait-группу для синхронизации работы горутин

	for _, value := range massiv {
		wgSquares.Add(1)
		go func(value int) { // Запуск горутины на вычисление квадратов

			sqr := value * value
			fmt.Println(sqr) // Считаем и выводим квадрат числа

			wgSquares.Done() // Завершаем горутину
		}(value)
	}

	wgSquares.Wait() // Ожидание заверешение работы всех запущенных горутин

}
