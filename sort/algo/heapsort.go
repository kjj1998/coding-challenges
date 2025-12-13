package algo

func maxHeapify(arr []string, heapSize int, index int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}
	if largest != index {
		arr[index], arr[largest] = arr[largest], arr[index]
		maxHeapify(arr, heapSize, largest)
	}
}

func buildHeap(arr []string) {
	heapSize := len(arr)
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(arr, heapSize, i)
	}
}

func HeapSort(arr []string) {
	heapSize := len(arr)
	buildHeap(arr)
	for i := heapSize - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapSize -= 1
		maxHeapify(arr, heapSize, 0)
	}
}
