package array

type PrefixSumResult[T Number] struct {
	arr []T
}

func (p *PrefixSumResult[T]) PrefixSumArray() []T {
	return p.arr
}
func (a *Array[T]) PrefixSum() *PrefixSumResult[T] {
	var sum_array []T
	var currentSum T = 0
	for _, val := range a.arr {
		currentSum = val + currentSum
		sum_array = append(sum_array, currentSum)
	}
	return &PrefixSumResult[T]{
		arr: sum_array,
	}
}
