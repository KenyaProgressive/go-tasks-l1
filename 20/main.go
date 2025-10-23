package main

import "fmt"

func main() {
	words := "East or west home is best"
	reversedWords := changeWordsOrder([]byte(words))

	fmt.Println("Before:", words)
	fmt.Println("After:", reversedWords)
}

func changeWordsOrder(s []byte) string { // Основная функция переворота порядка слов

	reverseString(s, 0, len(s)-1) // Переворачиваем строку

	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' { // Вычисляем границу подстроки -- конец всей строки или пробел
			reverseString(s, start, i-1) // Переворачиваем строку
			start = i + 1                // Двигаем начальную границу
		}
	}

	return string(s)

}

func reverseString(s []byte, start int, finish int) { // Функция переворота строки
	for start < finish { // Пока они не сравняются, меняем все символы местами
		s[start], s[finish] = s[finish], s[start]
		start++
		finish--
	}
}
