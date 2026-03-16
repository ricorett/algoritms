package selectioSort

func SelectionSort(arr []int) []int {
	count := len(arr)
	for i := 0; i < count; i++ {
		min := i                         //save minimum
		for j := i + 1; j < count; j++ { //next element from i
			if arr[j] < arr[min] { //if current less than minimum
				min = j // new min
			}

		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}
