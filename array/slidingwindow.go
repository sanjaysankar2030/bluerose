package array

import "fmt"

type SlidingWindowBuilder[T Number] struct {
	arr []T
}

func (array *Array[T]) SlidingWindow() *SlidingWindowBuilder[T] {
	return &SlidingWindowBuilder[T]{
		arr: array.arr,
	}
}

func (a *SlidingWindowBuilder[T]) FixedLength(k int) T {
	length := len(a.arr)
	if k >= length {
		fmt.Println("SubArray Out of Bounds.")
		return 0
	}
	var maxSum T
	for i := 0; i < k; i++ {
		maxSum = maxSum + a.arr[i]
	}
	windowSum := maxSum
	for i := k; i < length; i++ {
		windowSum = windowSum + a.arr[i] - a.arr[i-k]
		maxSum = max(windowSum, maxSum)
	}
	return maxSum
}

func (a *SlidingWindowBuilder[T]) VariableLength() T {
	length := len(a.arr)
	if length < 0 {
		return 0
	}
	var globalMaxSum T
	for k := 1; k < length; k++ {
		var maxSum T
		for i := 0; i < k; i++ {
			maxSum += a.arr[i]
		}
		windowMax := maxSum
		for i := k; i < length; i++ {
			maxSum = maxSum + a.arr[i] - a.arr[i-k]
			windowMax = max(maxSum, windowMax)
		}
		globalMaxSum = max(windowMax, globalMaxSum)
	}
	return globalMaxSum
}
