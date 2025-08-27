package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	workersNumber := setWorkersNumber() // Парсим количество запускаемых горутин
	if workersNumber <= 0 {
		fmt.Println("===Incorrect number of launching goroutines (-wn > 0)")
		return
	}
	myChannel := make(chan string) // Создаем небуферизрованный канал для "общения" горутин
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30) // Создаем тестовый контекст для автозавершения
	defer cancel()

	messageCount := 0 // Переменная с номером отправляемого сообщения

	wg.Add(workersNumber)

	for i := 0; i < workersNumber; i++ { // Запуск указанного кол-ва горутин для чтения сообщений из канала
		go func(ctx context.Context, ch chan string, workerId int) {
			for {
				select {
				case <-ctx.Done():
					wg.Done() // После окончания контекста горутина сама завершается
					return
				default:
					fmt.Printf("===GOT MESSAGE FROM CHANNEL BY WORKER %d:%s\n", workerId, <-ch)
					time.Sleep(time.Second * 1)
				}
			}
		}(ctx, myChannel, i)
	}

	for { // Запись сообщений в канал
		select {
		case <-ctx.Done():
			fmt.Printf("Task3 was succesfully finished\n")
			return
		default:
			msg := "message-with-code-" + strconv.Itoa(rand.Int()) // Создание сообщения
			myChannel <- msg
			fmt.Printf("===MESSAGE %d WAS PULLED IN CHANNEL\n", messageCount+1)
			messageCount++
			// Сделаем другую задержку, чтобы сообщений было заведомо больше, чем горутин
			time.Sleep(time.Millisecond * 800)
		}
	}

}

func setWorkersNumber() int { // Парсинг количества запускаемых горутин
	workersNumber := flag.Int("wn", 5, "Number of creating workers")
	flag.Parse()

	return *workersNumber
}
