package main

import "fmt"

func main() {

	before := "главрыба"
	wihoutExtraVarFlag := true
	after := reverseString(before, wihoutExtraVarFlag)

	fmt.Printf("%s ------> %s\n", before, after)
}

func reverseString(baseStrng string, wihoutExtraVarFlag bool) string {

	baseSymbolsSlice := []rune(baseStrng) // Конвертируем в срез рун -- так результат переворота будет точный
	countSymbols := len(baseSymbolsSlice) // Длина строки для расчёта индексов элементов с конца

	if wihoutExtraVarFlag {
		for i := 0; i < countSymbols/2; i++ {
			swapWithoutExtraVar(&baseSymbolsSlice[i], &baseSymbolsSlice[countSymbols-1-i]) // Меняем символы без третьей переменной
		}
	} else {
		for i := 0; i < countSymbols/2; i++ {
			symbolsSwap(&baseSymbolsSlice[i], &baseSymbolsSlice[countSymbols-1-i]) // Меняем последний символ с первым через третью переменную
		}
	}

	reversedString := string(baseSymbolsSlice) // Конвертируем результат в строку

	return reversedString

}

func symbolsSwap(a *rune, b *rune) { // Через указатели на элементы слайса меняем значения
	temp := *a
	*a = *b
	*b = temp
}

func swapWithoutExtraVar(a *rune, b *rune) { // Либо из 13 задачи без третьей переменной (т.к. rune это число -- а точнее номер символа в Unicode -- то сработает)

	*a += *b
	*b = (*b - *a) * (-1)
	*a -= *b

}
