package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Lab": {40.21321, 132.21312},
	"Google": {32.12312, 12.12312},
}

func main() {
	fmt.Println(m)
}
