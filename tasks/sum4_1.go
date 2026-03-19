package tasks

func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	mid := len(arr) / 2
	left := sum(arr[:mid])
	right := sum(arr[mid:])
	return left + right
}
