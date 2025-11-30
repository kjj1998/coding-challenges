package utils

import "slices"

func Sort(strs *[]string, unique bool) {
	slices.Sort(*strs)

	if unique {
		if len(*strs) < 2 {
			return
		}

		lastUniqueIndex := 0
		curIndex := 1

		for curIndex < len(*strs) {
			lastUnique := (*strs)[lastUniqueIndex]
			cur := (*strs)[curIndex]

			if cur != lastUnique {
				(*strs)[lastUniqueIndex+1] = cur
				lastUniqueIndex++
			}

			curIndex++
		}

		*strs = (*strs)[:lastUniqueIndex]
	}
}
