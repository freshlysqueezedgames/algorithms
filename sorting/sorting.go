package sorting

type Comparitor[K int] func(a K, b K) int

func SelectionSort[K int](list []K, comparitor Comparitor[K]) []K {
	return []K{}
}
