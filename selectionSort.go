package selectioSort

func SelectionSort(arr []int) []int {
	count := len(arr)
	for i := 0; i < count; i++ {
		minimum := i                     //save minimum
		for j := i + 1; j < count; j++ { //next element from i
			if arr[j] < arr[minimum] { //if current less than minimum
				minimum = j // new minimum
			}

		}
		if minimum != i {
			arr[i], arr[minimum] = arr[minimum], arr[i]
		}
	}
	return arr
}
