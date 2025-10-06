package main

func main() {

	slice1 := []int{2, 3, 4}
	slice2 := []int{3, 2, 5}

	set1 := toSet(slice1)
	set2 := toSet(slice2)

}

func toSet(my_slice []int) map[int]struct{} {

	result := make(map[int]struct{})

	for _, element := range my_slice {
		result[element] = struct{}{}
	}

	return result
}
