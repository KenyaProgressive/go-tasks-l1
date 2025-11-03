package main

import "fmt"

func main() {
	arr := []int{3, 4, 5, 6}
	newArr, err := deleteArrayElement(arr, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newArr)
}

func deleteArrayElement(arr []int, targetIndex int) ([]int, error) {
	if len(arr) == 0 { // Если длина 0, то выводим ошибку
		return nil, fmt.Errorf("слайс не существует или является пустым")
	} else { // [3, 4, 5, 6]
		copy(arr[targetIndex:], arr[targetIndex+1:]) // [3, 4, 6, 6] -- берем следующую за частью с удаляемым элементом часть слайса, и вставляем её на место 1-ой
		return arr[:len(arr)-1], nil                 // [3, 4, 6] -- Выводим n-1 элемент, чтобы было без повторений
	}

}
