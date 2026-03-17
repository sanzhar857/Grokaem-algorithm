package main

import "fmt"

func main() {
	// countdown(4)
	// fmt.Println(fact(5))
	fmt.Println(IsPowerOfTwo(3))
}

func countdown(num int) {
	fmt.Println(num)

	if num <= 1 {
		return
	} else {
		countdown(num - 1)
	}
}

func fact(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * fact(num-1)
	}
}
