package heapSort

func heapify(arr []int, n int, i int) {
	largest := i //инициализируем больший элемент как корень
	l := 2*i + 1 //левый
	r := 2*i + 2 // правый

	//если левый дочерний элемент больше корня
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	//если правый дочерний элемент больше, чем самый большой элемент на данный момент
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	// если самый большой элемент не корень
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		//рекурсивно преобразуем в двоичную кучу затронутое дерево
		heapify(arr, n, largest)
	}

}

func HeapSort(arr []int, n int) {
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}
