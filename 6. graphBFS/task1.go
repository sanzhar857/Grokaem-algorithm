package main

import "fmt"

func canReach[T comparable](graph Graph[T], start, target T) bool {

	if start == target {
		return true
	}

	queue := []T{start}
	visited := make(map[T]bool)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbour := range graph[current] {
			if !visited[neighbour] {
				visited[neighbour] = true
				if neighbour == target {
					return true
				}
			}
		}
	}
	return false
}

func cities() {
	kazakhstanRoads := Graph[string]{
		"Astana":    {"Karaganda", "Kokshetau"},
		"Karaganda": {"Astana", "Almaty", "Balkhash"},
		"Almaty":    {"Karaganda", "Shymkent", "Taldykorgan"},
		"Shymkent":  {"Almaty", "Taraz"},
		"Taraz":     {"Shymkent"},
		"Balkhash":  {"Karaganda"},
		"Kokshetau": {"Astana"},
		"Aktau":     {"Atyrau"}, // Опа! Этот кусок не связан с Астаной напрямую
		"Atyrau":    {"Aktau"},
	}

	fmt.Println(canReach(kazakhstanRoads, "Astana", "Karaganda"))
}
