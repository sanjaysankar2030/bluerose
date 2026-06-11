package main

import (
	"bluerose/array"
	"fmt"
)

func main() {
	arr := []int{10, 20, -40, 60, -50, 60}
	a := array.InitArray(arr)
	sum:=a.Kadane()
	fmt.Println(sum.WithSubarray())
	fmt.Println(sum)
	
//	prefixsum:=a.PrefixSum
//	sliding_window[]:=a.SlidingWindow
}

