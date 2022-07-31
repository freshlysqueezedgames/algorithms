package sorting

type Comparable interface {
	string | int | float32 | float64 | int32 | int64 | uint
}

// returns -1 for move a is less than b, 0 for equality / do nothing, 1 for a is more than b
type Comparitor[K Comparable] func(a K, b K) int

func SelectionSort[K Comparable](list []K, comparitor Comparitor[K]) {
	length := len(list)

	if length == 0 {
		return
	}

	lowestIndex := 0

	for i := 0; i < length; i++ {
		sorted := true

		for j := i; j < length; j++ {
			if comparitor(list[j], list[lowestIndex]) == -1 {
				sorted = false
				lowestIndex = j
			}
		}

		if sorted {
			return
		}

		temp := list[lowestIndex]
		list[lowestIndex] = list[i]
		list[i] = temp
	}
}
