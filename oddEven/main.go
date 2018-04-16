package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range numbers {
		fmt.Printf("%v is %v\n", num, evenOrOdd(num))
	}
}

func evenOrOdd(n int) string {
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}
