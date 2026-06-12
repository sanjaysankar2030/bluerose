package main

import (
	"bluerose/array"
	"fmt"
)

func main() {
	arr := []int{10, 20, 10, 5, 15}
	a := array.InitArray(arr)
//	fmt.Println(a.Length())
//	sum:=a.Kadane()
//	fmt.Println(sum.WithSubarray())
//	fmt.Println(sum)
prefixsum:=a.PrefixSum()
fmt.Print(prefixsum.PrefixSumArray())
//sliding_window[]:=a.SlidingWindow
}

