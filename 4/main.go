package main

/*
	Код взят и 3-ей задачи и будет модифицирован согласно условию задачи №4.
*/

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
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
	ctxSigInt, cacnelSigInt := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cacnelSigInt()

	messageCount := 0 // Переменная с номером отправляемого сообщения

	wg.Add(workersNumber)

	for i := 0; i < workersNumber; i++ { // Запуск указанного кол-ва горутин для чтения сообщений из канала
		go func(ctx context.Context, ch chan string, workerId int) {
			for {
				select {
				case <-ctxSigInt.Done(): // Если программа завершается через CTRL+C
					fmt.Print("\nGoroutine was finished by CTRL+C\n")
					wg.Done() // Завершение работы горутины
					return
				case msg, ok := <-myChannel: // Чтение из канала и обработка закрытия канала
					if !ok {
						wg.Done()
						return
					}
					fmt.Printf("===GOT MESSAGE FROM CHANNEL BY WORKER %d:%s\n", workerId, msg)
					time.Sleep(time.Second * 1)

				}
			}
		}(ctxSigInt, myChannel, i)
	}

	for { // Запись сообщений в канал
		select {
		case <-ctxSigInt.Done():
			fmt.Printf("\nTask4 was succesfully finished\n")
			close(myChannel)
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
