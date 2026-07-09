package main

import (
	"bluerose/array"
	"fmt"
)

func main() {
	arr := []int{10, 20, 10, 5, 15}
	// // arr2 := []int{5, 2, -1, 0, 3}
	// a := array.InitArray(arr)
	b := array.InitArray(arr)
	// fmt.Println(a.Length())
	// sum := b.Kadane()
	// fmt.Println(sum.WithSubarray())
	// fmt.Println("Kadane Sum :", sum)
	// prefixsum := a.PrefixSum()
	// fmt.Print(prefixsum.PrefixSumArray())
	// slidingWindow := b.SlidingWindow().FixedLength(3)
	// varslidingWindow := b.SlidingWindow().VariableLength()
	// fmt.Println("Sliding Window k == 3 :", slidingWindow)
	// fmt.Println("Variable Sliding Window :", varslidingWindow)
	keyMap := b.OppositeTwoPointer()
	fmt.Println(keyMap)
}
