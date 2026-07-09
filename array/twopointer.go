package array

type TwoPointerResult[T Number] struct {
	resultMap map[string][]T
}

func (a *Array[T]) OppositeTwoPointer() *TwoPointerResult[T] {
	left := 0
	right := len(a.arr) - 1
	result := map[string][]T{
		"left":  {},
		"right": {},
	}
	for i := 0; i <= right; i++ {
		result["left"] = append(result["left"], a.arr[left])
		result["right"] = append(result["right"], a.arr[right])
		right -= 1
		left += 1
	}
	return &TwoPointerResult[T]{
		resultMap: result,
	}
}
