package main

import "fmt"

func main() {
	c := make([]int, 3, 1)
	fmt.Println("%d,%d", len(c), cap(c))
}
