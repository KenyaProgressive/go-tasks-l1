package main

import (
	"context"
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	validValuesFlag := [5]string{"cond", "ctx", "hard-exit", "notify-channel", "close-channel"}
	isInvalidValueStop := 1  // 1 -- failure, 0 -- success
	stopFlag := kindOfStop() // Парсинг флага

	for _, value := range validValuesFlag {
		if stopFlag == value {
			isInvalidValueStop = 0 // Если значение в массиве корректных флагов -- Завершение проверки
			break
		}
	}

	if isInvalidValueStop == 1 {
		fmt.Println("Incorrect flag value") // Если значение флага некорректно -- завершение программы
		return
	}

	wg := sync.WaitGroup{}

	stopChannel := make(chan bool)
	switch stopFlag {
	case "cond": // Завершение по условию
		conditionExitFromGoroutine(&wg)
		return
	case "ctx": // Завершение по отмене контекста
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		contextExitFromGoroutine(&wg, ctx)
		return
	case "hard-exit": // Завершение через runtime.Goexit()
		runtimeGoExitFromGoroutine(&wg)
		return
	case "notify-channel": // Завершение через оповещение в канале
		signalStopToChannelExitFromGoroutine(&wg, stopChannel)
		return
	case "close-channel": // Завершение через закрытие канала
		closeChannelExitFromGoroutine(&wg, stopChannel)
		return

	}
}

//===============================================================================

func conditionExitFromGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)
	number := 0
	go func(num int, wg *sync.WaitGroup) {
		for {
			num++
			sum, count := sumAndCountDigits(num)
			fmt.Printf("sum: %d, count: %d\n", sum, count)
			time.Sleep(time.Second * 1)
			if sum == 6 || count == 3 {
				wg.Done()
				fmt.Println("Goroutine finished by completing condition")
				return
			}
		}
	}(number, wg)
	wg.Wait()
}

func contextExitFromGoroutine(wg *sync.WaitGroup, ctx context.Context) {
	wg.Add(1)
	number := 0
	go func(num int, ctx context.Context, wg *sync.WaitGroup) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine was finished by cancelling context")
				wg.Done()
				return
			default:
				num++
				fmt.Printf("Num: %d\n", num)
				time.Sleep(time.Second * 1)
			}
		}
	}(number, ctx, wg)

	wg.Wait()
}

func runtimeGoExitFromGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		defer fmt.Println("Goroutine was finished by runtime.Goexit()")
		for {
			fmt.Println("Goroutine was laucnhed . . .")
			time.Sleep(time.Second * 3) // Симуляция работы горутины
			runtime.Goexit()
		}

	}()
}

func signalStopToChannelExitFromGoroutine(wg *sync.WaitGroup, stopChannel chan bool) {
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer fmt.Println("Goroutine was finished by signal Stop in channel")
		for {
			select {
			case s := <-stopChannel:
				if s {
					wg.Done()
					return
				}
			default:
				fmt.Println("Working...")
				time.Sleep(time.Millisecond * 300)

			}
		}
	}(wg)

	time.Sleep(time.Second * 5)
	stopChannel <- true
	wg.Wait()
}

func closeChannelExitFromGoroutine(wg *sync.WaitGroup, stopChannel chan bool) {
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer fmt.Println("Goroutine was finished by closing channel")
		for {
			select {
			case <-stopChannel: // Если канал закрыт, вернется нулевое значение, и горутина завершится
				wg.Done()
				return
			default:
				fmt.Println("Working...")
				time.Sleep(time.Millisecond * 300)
			}
		}
	}(wg)

	time.Sleep(time.Second * 5)
	close(stopChannel)
	wg.Wait()
}

// ===============================================================================

func sumAndCountDigits(num int) (int, int) { // Сумма цифр числа и подсчет количества цифр в числе
	summ := 0
	countDigits := 0
	for num != 0 {
		summ += num % 10
		num /= 10
		countDigits += 1
	}
	return summ, countDigits
}

func kindOfStop() string { // Создание и парсинг флага
	s := flag.String("stop", "cond", "How would be stoped the goroutine")
	flag.Parse()
	return *s
}
