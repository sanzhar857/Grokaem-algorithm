package main

func fib(n int) int {
	if n == 1 {
		return 1
	}

	if n == 0 {
		return 0
	}

	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	memo := make(map[int]int)
	return fibMemo(memo, n)
}

func fibMemo(memo map[int]int, n int) int {
	if n <= 1 {
		return n
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = fibMemo(memo, n-1) + fibMemo(memo, n-2)
	return memo[n]
}
