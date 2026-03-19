package tasks

func countElems(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + countElems(arr[1:])

}

//[1 2 3 4]
// 1 + count([2 3 4])
// 1 + 1 + count([3 4])
// 1 + 1 + 1 + count([4])
// 1 + 1 + 1 + 1
