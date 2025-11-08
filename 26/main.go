package main

import (
	"fmt"
	"strings"
)

func main() {
	checkString1 := "abCdefAaf"
	checkString2 := "abcd"
	checkString3 := "aabbccdd"

	fmt.Println("Result1:", isStringWithUniqueSymbols(checkString1))
	fmt.Println("Result2:", isStringWithUniqueSymbols(checkString2))
	fmt.Println("Result3:", isStringWithUniqueSymbols(checkString3))

}

func isStringWithUniqueSymbols(s string) bool {
	lowerS := strings.ToLower(s)

	setS := makeSetFromLetters([]rune(lowerS)) // Сделать множество из слайса рун -- таким образом оставим только уникальные

	if len(lowerS) == len(setS) { // Если количество уникальных символов и длина строки равны -- возвращаем True
		return true
	} else {
		return false // Иначе символы повторяются -- возвращаем false
	}
}

func makeSetFromLetters(data []rune) map[rune]struct{} {
	result := make(map[rune]struct{})

	for _, letter := range data {
		if _, exist := result[letter]; !exist { // Если значение по ключу в мапе отсутствует, значит добавляем
			result[letter] = struct{}{}
		}
	}

	return result
}
