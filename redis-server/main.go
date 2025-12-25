package main

import "fmt"

func main() {
	simpleRedisString := "+OK\r\n"

	if len(simpleRedisString) == 0 {
		fmt.Println("no redis command detected")
	}
}
