package sorting

type Comparable interface {
	string | int | float32 | float64 | int32 | int64 | uint
}

// returns -1 for move a is less than b, 0 for equality / do nothing, 1 for a is more than b
type Comparitor[K Comparable] func(a, b K) int

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

		temp := list[lowestIndex]
		list[lowestIndex] = list[i]
		list[i] = temp
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
				temp := list[j]
				list[j] = list[i]
				list[i] = temp

				swapped = true
			}
		}

		if !swapped {
			return
		}
	}
}
