//reverse a string

package main

import "fmt"

func main() {
	input := "Hello"
	reversed := reverseString(input)
	fmt.Println(reversed)
}

func reverseString(input string) string {

	var reversed string
	for i := len(input) - 1; i >= 0; i-- {
		reversed = reversed + input[i:i+1]
	}
	return reversed
}
