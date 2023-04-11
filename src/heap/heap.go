package main

import (
	"container/heap"
	"fmt"
)

type Graph map[string]map[string]int

type Item struct {
	node     string
	distance int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func Dijkstra(graph Graph, start, end string) ([]string, int) {
	distances := make(map[string]int)
	for node := range graph {
		distances[node] = -1
	}
	distances[start] = 0
	previousNodes := make(map[string]string)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{start, 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		if item.node == end {
			path := []string{}
			node := end
			for node != start {
				path = append([]string{node}, path...)
				node = previousNodes[node]
			}
			path = append([]string{start}, path...)
			return path, item.distance
		}

		for neighbor, weight := range graph[item.node] {
			distance := item.distance + weight
			if distances[neighbor] == -1 || distance < distances[neighbor] {
				distances[neighbor] = distance
				previousNodes[neighbor] = item.node
				heap.Push(&pq, &Item{neighbor, distance})
			}
		}
	}

	return []string{}, -1
}

func main() {
	graph := Graph{
		"A": {"B": 7, "C": 9, "F": 14},
		"B": {"A": 7, "C": 10},
	}

	fmt.Println(Dijkstra(graph, "A", "B"))
}
