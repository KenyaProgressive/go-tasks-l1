package main

import (
	"fmt"
	"sync"
)

type ConcurrentCounter struct {
	counterValue int
	mu           sync.Mutex
}

func (cc *ConcurrentCounter) Increase() {
	cc.mu.Lock() // Захват мьютекса для инкрементации
	cc.counterValue++
	cc.mu.Unlock() // Мьютекс отдаётся, следующая горутина начинает выполняться

}

func main() {
	wg := sync.WaitGroup{}
	cc := ConcurrentCounter{counterValue: 0, mu: sync.Mutex{}}

	for i := 0; i < 18000; i++ {
		wg.Add(1)
		go func(cc *ConcurrentCounter) {
			defer wg.Done()
			cc.Increase()
		}(&cc)
	}

	wg.Wait() // Дожидаемся пока все горутины завершатся, чтобы получить правильное значение в результате

	fmt.Println("Result:", cc.counterValue)
}
