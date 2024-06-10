package main

import "fmt"

func compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	}
	return 1
}

func main() {
	fmt.Println(compare("hello", "hello"))
	fmt.Println(compare("alot", "lot"))
	fmt.Println(compare("hey", "ben"))
}
