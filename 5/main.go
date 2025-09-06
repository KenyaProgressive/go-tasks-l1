package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	ch := make(chan int)

	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()

	fmt.Println("Programm was finifhed by ending time")

}

func producer(ch chan int, wg *sync.WaitGroup) {
	t := time.After(time.Second * 15) // Таймер на 15 секунд
	defer wg.Done()
	for {
		msg := rand.Int()
		select {
		case ch <- msg: // Если канал открыт -- сообщение отправляется
			fmt.Printf("sent message: %d\n", msg)
			time.Sleep(time.Millisecond * 800)
		case <-t: // По истечению времени, канал закрывается
			close(ch)
			return
		}
	}
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch { // Чтение до момента закрытия канала
		fmt.Printf("message was recieved: %d\n", msg)
		time.Sleep(time.Millisecond * 800)
	}
}
