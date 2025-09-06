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
	t := time.After(time.Second * 15)
	defer wg.Done()
	for {
		msg := rand.Int()
		select {
		case ch <- msg:
			fmt.Printf("sent message: %d\n", msg)
			time.Sleep(time.Millisecond * 800)
		case <-t:
			close(ch)
			return
		}
	}
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch {
		fmt.Printf("message was recieved: %d\n", msg)
		time.Sleep(time.Millisecond * 800)
	}
}
