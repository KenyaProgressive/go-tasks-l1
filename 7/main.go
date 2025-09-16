package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	mapType := parseFlagMapType() // Парсинг флага вида мапы

	switch mapType {
	case "usual": // Запись классическую мапу
		myMap := make(map[int]string)
		concWriteToUsualMap(myMap, &wg)
		fmt.Println("\nProgramm was successfully finished")
		return
	case "conc": // Запись в sync.Map
		myMapSync := sync.Map{}
		concWriteToConcMap(&myMapSync, &wg)
		fmt.Println("\nProgramm was successfully finished")
		return
	default:
		fmt.Printf("Incorrect flag value: %s. Choose between usual and conc.\n", mapType)
		return
	}

}

func concWriteToUsualMap(myMap map[int]string, wg *sync.WaitGroup) {
	mute := sync.Mutex{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(myMap map[int]string, wg *sync.WaitGroup, mute *sync.Mutex, index int) {
			defer wg.Done()

			myValue := fmt.Sprintf("key-with-index-%d", index) // Формируем значение

			mute.Lock() // Блокируем другие горутины

			myMap[index] = myValue // Выполняем запись в мапу

			mute.Unlock() // Снимаем блокировку

			time.Sleep(time.Microsecond * 300)

		}(myMap, wg, &mute, i)
	}

	wg.Wait()

	fmt.Println("Saved values ===>") // Печать сохраненных значений
	for key := range myMap {
		fmt.Printf("key: %d; value: %s\n", key, myMap[key])
		time.Sleep(time.Millisecond * 500)
	}

}

func concWriteToConcMap(myMapSync *sync.Map, wg *sync.WaitGroup) {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(myMapSync *sync.Map, index int, wg *sync.WaitGroup) {
			defer wg.Done()

			myValue := fmt.Sprintf("key-with-index-%d", index)

			myMapSync.Store(index, myValue) // Конкуретная запись в мапу

			time.Sleep(time.Microsecond * 300)
		}(myMapSync, i, wg)
	}

	wg.Wait()

	fmt.Println("Saved values ===>") // Печать сохраненных значений
	myMapSync.Range(func(key, value any) bool {
		fmt.Printf("key: %d; value: %s\n", key, value)
		time.Sleep(time.Millisecond * 500)
		return true
	})
}

func parseFlagMapType() string {
	mapType := flag.String("map-type", "usual", "Type of map -- usual or conc")
	flag.Parse()
	return *mapType
}
