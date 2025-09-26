package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	numCh := make(chan int)        // Канал для отправки чисел
	numDoubledCh := make(chan int) // Канал для чтения удвоенных чисел
	numsSlice := make([]int, 0)

	var countNums int
	if err := readCountNumbers(&countNums); err != nil {
		return
	} // Введем количество чисел

	numGenerator(&numsSlice, countNums) // Сгенерируем и вставим числа в слайс

	wg.Add(2)

	go numsPusher(&numsSlice, numCh, &wg)     // Запись чисел в канал №1
	go numsDoubling(numCh, numDoubledCh, &wg) // Чтение чисел, удваивание и запись в канал №2

	for value := range numDoubledCh { // Вывод чисел на экран, пока канал не пуст и не закрыт
		fmt.Printf("==> got modified value from ch2: %d\n", value)
		countNums--
	}

	wg.Wait()

	fmt.Println("Programm was successfully finished")

}

func numsDoubling(firstCh chan int, secondCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range firstCh {
		modifiedNum := value * 2
		secondCh <- modifiedNum
		time.Sleep(time.Millisecond * 90)
	}

	close(secondCh)

}

func numsPusher(nums *[]int, firstCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for index, num := range *nums {

		firstCh <- num
		fmt.Printf("==> Num %d with index %d was successfully pushed\n", num, index)
		time.Sleep(90 * time.Millisecond)
	}

	close(firstCh)
}

func numGenerator(nums *[]int, countNums int) {
	seed := rand.NewSource(time.Now().UnixNano()) // Создаем уникальное стартовое значение для генератора чисел
	generator := rand.New(seed)                   // Генератор на основе сида

	for len(*nums) < countNums {
		*nums = append(*nums, generator.Intn(999))
		time.Sleep(time.Millisecond * 60)
	}
	fmt.Println("==> Nums are generated.")

}

func readCountNumbers(countNumber *int) error {
	scanner := bufio.NewScanner(os.Stdin)
	var err error

	fmt.Print("==> Quantity of numbers: ")
	scanner.Scan()

	*countNumber, err = strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Entered value is not an int number")
		return err
	}

	return nil

}
