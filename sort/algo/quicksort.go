package algo

func QuickSort(strs []string) []string {
	var less, more, pivotList []string

	if len(strs) <= 1 {
		return strs
	} else {
		pivot := medianOfThree(strs[0], strs[len(strs)-1], strs[len(strs)/2])
		for _, str := range strs {
			if str < pivot {
				less = append(less, str)
			} else if str > pivot {
				more = append(more, str)
			} else {
				pivotList = append(pivotList, str)
			}
		}

		less := QuickSort(less)
		more := QuickSort(more)

		result := append(append(less, pivotList...), more...)
		return result
	}
}

func medianOfThree(a, b, c string) string {
	if (a >= b && a <= c) || (a >= c && a <= b) {
		return a
	} else if (b >= a && b <= c) || (b >= c && b <= a) {
		return b
	}
	return c
}
