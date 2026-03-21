package main

import "fmt"

func main() {
	// words := []string{"apple", "orange", "apple", "banana", "apple", "orange"}
	// result := wordCounter(words)
	// fmt.Println(result)

	// fmt.Println(isAnagram("silent1", "listen2"))

	fmt.Println(twoSum([]int{2, 7, 11, 5}, 9))
}

func base() {
	m := make(map[string]int)

	m["Aktobe"] = 4
	m["Astana"] = 1

	val, ok := m["Aktobe"]
	if ok {
		fmt.Println(val)
	}
}

// Условие: Напиши функцию, которая принимает срез строк (например, названия городов или фруктов)
// и возвращает map, где ключи — это названия, а значения — сколько раз это название встретилось в списке.
func wordCounter(values []string) map[string]int {
	m := make(map[string]int, len(values))

	for _, val := range values {
		m[val]++
	}

	return m
}

// Условие: Даны две строки. Нужно определить, являются ли они анаграммами
// (состоят ли они из одних и тех же букв в одинаковом количестве, например: "silent" и "listen").
func isAnagram(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	m := make(map[rune]int)

	for _, char := range w1 {
		m[char]++
	}

	for _, char := range w2 {
		m[char]--
	}

	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}
