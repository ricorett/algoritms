package quickSort

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := len(arr) / 2

	lessArr := make([]int, pivot)
	moreArr := make([]int, pivot)

	for i := 0; i < len(arr); i++ {
		if i == pivot {
			continue
		}
		if arr[i] < arr[pivot] {
			lessArr[i] = arr[i]
		}
		if arr[i] > arr[pivot] {
			moreArr[i] = arr[i]
		}
	}
	combined := append(QuickSort(lessArr), QuickSort(moreArr)...)

	return combined
}
