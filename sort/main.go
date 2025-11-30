package main

import (
	"flag"
	"fmt"

	utils "github.com/kjj1998/coding-challenges/sort/utils"
)

func main() {
	flag.Parse()

	var filePath string
	if len(flag.Args()) > 0 {
		filePath = flag.Args()[0]
	}

	strs := utils.Read(&filePath)
	utils.Sort(strs)

	for _, str := range *strs {
		fmt.Println(str)
	}
}
