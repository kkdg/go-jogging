package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell labs"] = Vertex{
		40.1231, 32.1231,
	}
	fmt.Println(m["Bell labs"])
}
