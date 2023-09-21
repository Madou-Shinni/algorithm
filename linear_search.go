package algorithm

type LinearSearch[T comparable] struct {
}

func (LinearSearch[T]) Search(data []T, target T) int {
	for i, item := range data {
		if item == target {
			return i
		}
	}

	return -1
}
