package main

import (
	"log"

	"github.com/freshlysqueezedgames/algorithms/sorting"
)

func comparitor(a int, b int) int {
	if a < b {
		return -1
	}

	if a > b {
		return 1
	}

	return 0
}

func main() {
	log.Printf("sorted: %v", sorting.SelectionSort([]int{64, 25, 15, 22, 11}, comparitor))
}
