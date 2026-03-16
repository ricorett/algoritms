package insertionSort

func InsertionSort(arr []int) []int {
	count := len(arr)
	for i := 1; i < count; i++ { //from 1 to end
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- { //from i, j bigger than 0, and previous
			//element bigger than current, j--
			arr[j-1], arr[j] = arr[j], arr[j-1] //switch places
		}
	}
	return arr
}
