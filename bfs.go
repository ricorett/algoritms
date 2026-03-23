package bfs

import "fmt"

func bfs(graph map[int][]int, start int) {

	queue := []int{start}
	visited := map[int]bool{start: true}

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		fmt.Println(node)

		for _, neighbor := range graph[node] {

			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}
