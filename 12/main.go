package main

import "fmt"

func main() {

	data := []string{"cat", "cat", "dog", "cat", "tree"}

	set := makeSet(data)

	func(set map[string]struct{}) {
		counter := 0
		fmt.Print("Result: {")
		for key := range set {
			fmt.Printf("%s", key)
			counter++
			if counter != len(set) { // Постановка запятых между числами до последнего значения не включительно
				fmt.Print(", ")
			}
		}
		fmt.Println("}")
	}(set)

}

func makeSet(data []string) map[string]struct{} {
	result := make(map[string]struct{})

	for _, value := range data {
		if _, exist := result[value]; !exist { // Если значение по данному ключу есть в другом слайсе, значит ключ попадает в мапу-результат
			result[value] = struct{}{}
		}
	}

	return result
}
