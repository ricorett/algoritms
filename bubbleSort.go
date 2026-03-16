package bubbleSort

func BubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ { // iterate from begin
		for j := 0; j < length-i-1; j++ { //iterate from end
			changes := false
			if arr[j] > arr[j+1] { //if current bigger than next
				arr[j], arr[j+1] = arr[j+1], arr[j] //switch places
				changes = true
			}
			if !changes { //if no changes just return
				return arr
			}
		}
	}
	return arr
}
