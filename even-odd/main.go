package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range n {
		s := strconv.Itoa(num)

		if num%2 == 0 {
			fmt.Println(s + " is even")
		} else {
			fmt.Println(s + " is odd")
		}
	}
}
