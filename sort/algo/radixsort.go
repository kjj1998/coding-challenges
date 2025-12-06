package algo

func RadixSort(strs []string) []string {
	var maxStringLength = 0

	for _, str := range strs {
		maxStringLength = max(maxStringLength, len(str))
	}

	for i := maxStringLength - 1; i >= 0; i-- {
		strs = countingSort(strs, i)
	}

	return strs
}

func countingSort(strs []string, position int) []string {
	var tempArr [257]int

	for _, str := range strs {
		var charValue int
		if position >= len(str) {
			charValue = 0
		} else {
			charValue = int(str[position])
		}
		tempArr[charValue]++
	}

	for i := 1; i < len(tempArr); i++ {
		tempArr[i] = tempArr[i-1] + tempArr[i]
	}

	sortedOutput := make([]string, len(strs))

	for i := len(strs) - 1; i >= 0; i-- {
		str := strs[i]
		var charValue int
		if position >= len(str) {
			charValue = 0
		} else {
			charValue = int(str[position])
		}
		sortedOutputIndexPosition := tempArr[charValue] - 1
		tempArr[charValue]--

		sortedOutput[sortedOutputIndexPosition] = str
	}

	return sortedOutput
}
