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

func (a *SlidingWindowBuilder[T]) FixedLength(window_length int) T {
	length := len(a.arr)
	if window_length >= length {
		fmt.Println("SubArray Out of Bounds.")
		return 0
	}
	var maxSum T
	for i := 0; i < window_length; i++ {
		maxSum = maxSum + a.arr[i]
	}
	windowSum := maxSum
	for i := window_length; i < length; i++ {
		windowSum = windowSum + a.arr[i] - a.arr[i-window_length]
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
	for window_length := 1; window_length < length; window_length++ {
		var maxSum T
		for i := 0; i < window_length; i++ {
			maxSum += a.arr[i]
		}
		windowMax := maxSum
		for i := window_length; i < length; i++ {
			maxSum = maxSum + a.arr[i] - a.arr[i-window_length]
			windowMax = max(maxSum, windowMax)
		}
		globalMaxSum = max(windowMax, globalMaxSum)
	}
	return globalMaxSum
}
