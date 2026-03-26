package main

import (
	"fmt"
)

type Tree[T comparable] map[T][]T

func main() {
	names := Tree[string]{
		"Pics": {"2001", "odyssey"},
		"2001": {"a.png", "space.png"},
	}

	printNames(names, "Pics")
}

// обход всех элементов
func printNames[T string](graph Tree[T], startNode T) {
	queue := []T{startNode}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		// сюда можно добавить проверку
		fmt.Println(currentNode)

		for _, nei := range graph[currentNode] {
			queue = append(queue, nei)
		}

	}
}
