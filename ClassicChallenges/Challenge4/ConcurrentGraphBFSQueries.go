package main

import (
	"fmt"
	"sync"
)

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {2, 3},
		2: {3},
		3: {4},
		4: {},
	}
	queries := []int{0, 1, 2}
	numWorkers := 2

	result := ConcurrentBFSQueries(graph, queries, numWorkers)

	for _, query := range queries {
		fmt.Printf("BFS from %d: %v\n", query, result[query])
	}

}

func ConcurrentBFSQueries(graph map[int][]int, queries []int, numWorkers int) map[int][]int {
	if numWorkers <= 0 {
		return map[int][]int{}
	}

	jobChan := make(chan int, 2*numWorkers)
	resultChan := make(chan BFSResult, 2*numWorkers)
	result := map[int][]int{}
	wg := sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Workers(graph, jobChan, resultChan, &wg)
	}

	go func() {
		for _, query := range queries {
			jobChan <- query
		}
		close(jobChan)
	}()

	go func(resultChan chan BFSResult, wg *sync.WaitGroup) {
		wg.Wait()
		close(resultChan)
	}(resultChan, &wg)

	for traversal := range resultChan {
		result[traversal.Node] = traversal.BFSSearch
	}
	return result
}

type BFSResult struct {
	Node      int
	BFSSearch []int
}

func Workers(graph map[int][]int, jobChan chan int, resultChan chan BFSResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobChan {
		traversal := BFS(graph, job)
		resultChan <- BFSResult{job, traversal}
	}

}

func BFS(graph map[int][]int, startNode int) []int {
	result := []int{}
	queue := []int{}
	queue = append(queue, startNode)
	visited := map[int]bool{}
	visited[startNode] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		neighbours, exist := graph[current]
		if exist {
			for _, neighbour := range neighbours {
				if !visited[neighbour] {
					visited[neighbour] = true
					queue = append(queue, neighbour)
				}
			}
		}

	}
	return result
}
