package main

import "fmt"

func main() {
	// countdown(4)
	// fmt.Println(fact(5))
	// fmt.Println(IsPowerOfTwo(3))
	// fmt.Println(isPowerOfThree(0))
	// fmt.Println(isPowerOfFour(1024))
	fmt.Println(fib(10))
	fmt.Println(fib2(10))
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
