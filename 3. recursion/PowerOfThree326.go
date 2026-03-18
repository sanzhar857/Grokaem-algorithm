package main

func isPowerOfThree(n int) bool {

	if (n%3 == 0 || n == 1) && n != 0 {
		if n == 1 {
			return true
		} else {
			return isPowerOfThree(n / 3)
		}
	} else {
		return false
	}
}
