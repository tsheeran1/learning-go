package main

import "fmt"

func main() {

	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range ints {
		if num%2 == 0 {
			fmt.Printf("%v is Even\n", num)
		} else {
			fmt.Printf("%v is Odd\n", num)
		}
	}
}
