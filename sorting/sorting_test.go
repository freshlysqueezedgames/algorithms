package sorting

import (
	"fmt"
	"log"
	"testing"
	"time"
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

func track(name string) func() {
	now := time.Now()

	return func() {
		log.Printf("%v: %v", name, time.Since(now).String())
	}
}

func Test_SelectionSortSimpleArray(t *testing.T) {
	defer track("Total Time")()
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
	defer track("Total Time")()
	test := []int{64, 25, 15, 22, 11}

	BubbleSort(test, comparitor)

	for i, v := range []int{11, 15, 22, 25, 64} {
		if test[i] != v {
			t.Fatalf("Unexpect value %v at index: %v, expected %v (output %+v)", test[i], i, v, test)
		}
	}

	log.Printf("sorted: %+v", test)
}

func Test_InsertSortSimpleArray(t *testing.T) {
	defer track("Total Time")()

	test := []int{64, 25, 15, 22, 11}

	InsertionSort(test, comparitor)

	for i, v := range []int{11, 15, 22, 25, 64} {
		if test[i] != v {
			t.Fatalf("Unexpect value %v at index: %v, expected %v (output %+v)", test[i], i, v, test)
		}
	}

	fmt.Printf("sorted: %+v", test)
}
