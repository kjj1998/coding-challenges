package main

import (
	"fmt"

	"github.com/kjj1998/coding-challenges/sort/algo"
)

// "flag"

func main() {
	// unique := flag.Bool("u", false, "unique flag, determines whether sorted results should show unique values i.e. no duplicates")

	// flag.Parse()

	// var filePath string
	// if len(flag.Args()) > 0 {
	// 	filePath = flag.Args()[0]
	// }

	// strs := utils.Read(&filePath)
	// utils.Sort(strs, *unique)

	// for _, str := range *strs {
	// 	fmt.Println(str)
	// }

	strs := []string{"banana", "app", "apple", "bat"}

	strs = algo.RadixSort(strs)

	fmt.Println(strs)

}
