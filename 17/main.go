package main

import (
	"fmt"
)

type searchStatus struct {
	message string
}

func main() {
	sortedValues := []int{2, 4, 6, 8}

	index, status := binarySearch(sortedValues, 2)

	if status.message != "Success" {
		fmt.Println(status.message)
	} else {
		fmt.Println("Value was found and have index:", index)
	}
}

func binarySearch(sortedValues []int, target int) (int, searchStatus) {

}
