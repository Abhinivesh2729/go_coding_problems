//Check if an Array Contains Duplicates

package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 3}
	result := checkDuplicate(array)
	if result {
		fmt.Println("The array contains duplicates")
	} else {
		fmt.Println("No duplicates in the array")
	}
}

func checkDuplicate(array []int) bool {

	for _, number := range array {
		for i := 0; i <= len(array); i++ {
			if number == (array)[i] {
				return true
			}
		}
	}
	return false
}
