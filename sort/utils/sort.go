package utils

import (
	"slices"
)

func Sort(strs []string) []string {
	slices.Sort(strs)

	return strs
}

func Unique(results *[]string) {
	if len(*results) < 2 {
		return
	}

	lastUniqueIndex := 0
	curIndex := 1

	for curIndex < len(*results) {
		lastUnique := (*results)[lastUniqueIndex]
		cur := (*results)[curIndex]

		if cur != lastUnique {
			(*results)[lastUniqueIndex+1] = cur
			lastUniqueIndex++
		}

		curIndex++
	}

	*results = (*results)[:lastUniqueIndex+1]
}
