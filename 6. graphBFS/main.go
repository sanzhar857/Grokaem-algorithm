package main

import (
	"fmt"
)

type Graph[T comparable] map[T][]T

func BFS[T comparable](graph Graph[T], startNode T) {
	queue := []T{startNode}

	visited := make(map[T]bool)
	visited[startNode] = true

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		fmt.Printf("Посетили узел: %v\n", currentNode)

		for _, neighbour := range graph[currentNode] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}
}

func main() {
	socialGraph := Graph[string]{
		"Sanzhar": {"Asem", "Marjan"},
		"Asem":    {"Sanzhar", "Tumar"},
		"Marjan":  {"Sanzhar", "Astana"},
		"Tumar":   {"Asem"},
		"Astana":  {"Marjan"},
	}

	BFS(socialGraph, "Sanzhar")
}
