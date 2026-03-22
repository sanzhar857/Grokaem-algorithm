package main

import "fmt"

// type Graph[T comparable] map[T][]T

func persons() {
	socialGraph1 := Graph[string]{
		"Sanzhar": {"Asem", "Marjan"},
		"Asem":    {"Sanzhar", "Tumar", "Arman"},
		"Marjan":  {"Sanzhar", "Astana"},
		"Tumar":   {"Asem", "Astana"},
		"Arman":   {"Asem"},
		"Astana":  {"Marjan", "Tumar"},
	}

	fmt.Println(shortestPath(socialGraph1, "Sanzhar", "Astana"))
}

func shortestPath[T comparable](graph Graph[T], start, target T) int {

	if start == target {
		return 0
	}

	queue := []T{start}
	counter := make(map[T]int)
	counter[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbour := range graph[current] {
			if _, ok := counter[neighbour]; !ok {
				counter[neighbour] = counter[current] + 1
			}

			if neighbour == target {
				return counter[neighbour]
			}

			queue = append(queue, neighbour)
		}
	}
	return -1
}
