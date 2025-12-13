package algo

func MergeSort(strs []string) []string {
	length := len(strs)

	if length == 1 {
		return strs
	}

	midPoint := length / 2

	leftHalf := MergeSort(strs[:midPoint])
	rightHalf := MergeSort(strs[midPoint:])

	return merge(leftHalf, rightHalf)
}

func merge(left []string, right []string) []string {
	output := []string{}
	leftIndex, rightIndex := 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		leftStr := left[leftIndex]
		rightStr := right[rightIndex]

		if leftStr <= rightStr {
			output = append(output, leftStr)
			leftIndex++
		} else {
			output = append(output, rightStr)
			rightIndex++
		}
	}

	if leftIndex < len(left) {
		output = append(output, left[leftIndex:]...)
	} else if rightIndex < len(right) {
		output = append(output, right[rightIndex:]...)
	}

	return output
}
