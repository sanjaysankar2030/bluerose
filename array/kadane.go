package array

import "fmt"

type KadaneResult[T Number] struct {
	sum        T
	startIndex int
	endIndex   int
	arr        []T
}

func (r *KadaneResult[T]) Sum() T {return r.sum}
func (r *KadaneResult[T]) Start() int { return r.startIndex }
func (r *KadaneResult[T]) End() int   { return r.endIndex }
func (r *KadaneResult[T]) WithSubarray() ([]T, T) {
	return r.arr[r.startIndex : r.endIndex+1], r.sum
}

func (r *KadaneResult[T]) String() string {
	return fmt.Sprintf(
		"Sum: %v | Range: [%d, %d] | Subarray: %v",
		r.sum, r.startIndex, r.endIndex,
		r.arr[r.startIndex:r.endIndex+1],
	)
}

func (a *Array[T]) Kadane() *KadaneResult[T] {
	var current_sum T
	var global_sum T
	start, end, tempStart := 0, 0, 0

	for i, val := range a.arr {
		current_sum += val
		if current_sum > global_sum {
			global_sum = current_sum
			start = tempStart
			end = i
		}
		if current_sum < 0 {
			current_sum = 0
			tempStart = i + 1

		}
		if global_sum < 0 {
			global_sum = 0
		}
	}
	return &KadaneResult[T]{
		sum:        global_sum,
		startIndex: start,
		endIndex:   end,
		arr:        a.arr,
	}
}
