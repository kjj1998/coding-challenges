package main

import (
	"flag"
	"fmt"

	"github.com/kjj1998/coding-challenges/sort/algo"
	"github.com/kjj1998/coding-challenges/sort/utils"
)

func main() {
	unique := flag.Bool("u", false, "unique flag, determines whether sorted results should show unique values i.e. no duplicates")
	radixSort := flag.Bool("radixsort", false, "radix sort flag, determines whether results should be sorted using the radix sort algorithm")
	mergeSort := flag.Bool("mergesort", false, "merge sort flag, determines whether results should be sorted using the merge sort algorithm")
	quickSort := flag.Bool("quicksort", false, "quick sort flag, determines whether results should be sorted using the quick sort algorithm")
	heapSort := flag.Bool("heapsort", false, "heap sort flag, determines whether results should be sorted using the heap sort algorithm")

	flag.Parse()

	var filePath string
	if len(flag.Args()) > 0 {
		filePath = flag.Args()[0]
	}

	strs := utils.Read(&filePath)
	var results []string

	if *radixSort {
		results = algo.RadixSort(*strs)
	} else if *mergeSort {
		results = algo.MergeSort(*strs)
	} else if *quickSort {
		results = algo.QuickSort(*strs)
	} else if *heapSort {
		algo.HeapSort(*strs)
		results = *strs
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
