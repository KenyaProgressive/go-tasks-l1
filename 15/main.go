package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Проблемы кода ниже:

/* 1. Глобальная переменная -- плохое решение, потому что её изменение может привести к одной функции, может привести к поломке другой,
где она также используется.
2. Слайс, по сути, под капотом является структурой, хранящий указатель на первый элемент слайса в оригинальном массиве, длину и ёмкость. Поэтому,
при записи в переменную слайса от исходной строки заблокирует очистку GC исходной строки, потому что переменная со слайсом будет ссылаться на изначальную строку.
3. Строка в Go -- массив байтов. Соответственно, срезы от строк получаются через отрез количества байт, а не символов, что может привести к тому, что символ может оказаться "разрезанным".
4. Если длина строки меньше 100, то прозойдет panic -- out of range
*/

// var justString string

// func someFunc() {
//   v := createHugeString(1 &lt;&lt; 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

func main() {
	result := someFunc()
	fmt.Println("Результат-строка: ", someFunc())
	fmt.Printf("Длина строки (в байтах): %d;\n Количество символов: %d\n", len(result), utf8.RuneCountInString(result))
}

func someFunc() string {

	v := []rune(createHugeString(1 << 10)) // Строка разбивается на руны (занимает больше памяти, но при этом безопасно делиться по символам)

	sliceLength := 100

	if len(v) < sliceLength { // Если длина меньше, чем 100 -- отрезаем по длину строки
		sliceLength = len(v)
	}

	justString := string(v[:sliceLength])

	return justString

}

func createHugeString(length int) string {
	return strings.Repeat("HugeString1", length)
}
