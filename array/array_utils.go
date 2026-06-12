package array

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float // covers ALL int + float types
}
type Array[T Number] struct {
	arr []T // array of type T elements
}

// Takes a Generic slice and initialized the Array struct
func InitArray[T Number](values []T) *Array[T] {
	return &Array[T]{
		arr: values,
	}
}

// Prints the Array
func (a *Array[T]) PrintArray() {
	fmt.Println(a.arr)
}

func (a *Array[T]) Length() int {
	return len(a.arr)
}
