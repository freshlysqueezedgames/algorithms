package sorting

type Comparable interface {
	string | int | float32 | float64 | int32 | int64 | uint
}

// returns -1 for move a is less than b, 0 for equality / do nothing, 1 for a is more than b
type Comparitor[K Comparable] func(a, b K) int

func swap[K Comparable](list []K, i, j int) {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}

func SelectionSort[K Comparable](list []K, comparitor Comparitor[K]) {
	length := len(list)

	if length == 0 {
		return
	}

	for i := 0; i < length-1; i++ {
		sorted := true
		lowestIndex := i

		for j := i + 1; j < length; j++ {
			if comparitor(list[j], list[lowestIndex]) == -1 {
				sorted = false
				lowestIndex = j
			}
		}

		if sorted {
			return
		}

		swap(list, lowestIndex, i)
	}
}

func BubbleSort[K Comparable](list []K, comparitor Comparitor[K]) {
	length := len(list)

	if length == 0 {
		return
	}

	for {
		swapped := false

		for j := 0; j < length-1; j++ {
			var i int = j + 1

			if comparitor(list[i], list[j]) == -1 {
				swap(list, j, i)
				swapped = true
			}
		}

		if !swapped {
			return
		}
	}
}

// InsertionSort: Uses insertion sort to order items in a list, given a comparison function
func InsertionSort[K Comparable](list []K, comparitor Comparitor[K]) {
	length := len(list)

	if length == 0 {
		return
	}

	for j := 0; j < length-1; j++ {
		var i int = j + 1

		if comparitor(list[i], list[j]) == -1 {
			swap(list, j, i)

			if j > 0 {
				j -= 2 // drop back an extra index as we will be pulled forward 1 in the next loop
			}
		}
	}
}
