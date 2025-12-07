package main

import (
	"flag"
	"fmt"

	"github.com/kjj1998/coding-challenges/sort/algo"
	"github.com/kjj1998/coding-challenges/sort/utils"
)

func main() {
	unique := flag.Bool("u", false, "unique flag, determines whether sorted results should show unique values i.e. no duplicates")
	radixSort := flag.Bool("radixsort", false, "radixsort flag, determines whether results should be sorted using the radix sort algorithm")

	flag.Parse()

	var filePath string
	if len(flag.Args()) > 0 {
		filePath = flag.Args()[0]
	}

	strs := utils.Read(&filePath)

	// testStr := []string{"banana", "app", "apple", "bat"}
	var results []string

	if *radixSort {
		results = algo.RadixSort(*strs)
	} else {
		results = utils.Sort(*strs)
	}

	if *unique {
		utils.Unique(&results)
	}

	for _, str := range results {
		fmt.Println(str)
	}
}
