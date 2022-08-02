package sorting

import (
	"log"
	"testing"
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

func Test_SelectionSortSimpleArray(t *testing.T) {
	test := []int{64, 25, 15, 22, 11}

	SelectionSort(test, comparitor)

	for i, v := range []int{11, 15, 22, 25, 64} {
		if test[i] != v {
			t.Fatalf("Unexpect value %v at index: %v, expected %v", test[i], i, v)
		}
	}

	log.Printf("sorted: %+v", test)
}

func Test_BubbleSortSimpleArray(t *testing.T) {
	test := []int{64, 25, 15, 22, 11}

	BubbleSort(test, comparitor)

	for i, v := range []int{11, 15, 22, 25, 64} {
		if test[i] != v {
			t.Fatalf("Unexpect value %v at index: %v, expected %v (output %+v)", test[i], i, v, test)
		}
	}

	log.Printf("sorted: %+v", test)
}
