// check if a number is even
package main

import "fmt"

func main() {
	if checkEvenNumber(4) {
		fmt.Println("The number is Even")
	} else {
		fmt.Println("The number is Odd")
	}
}

func checkEvenNumber(number int) bool {
	if number/2 == 0 {
		return true
	}
	return false
}
