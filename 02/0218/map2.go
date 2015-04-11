package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.42232, -12.132323,
	},
	"Google": Vertex{
		31.2314, -122.1313,
	},
}

func main() {
	fmt.Println(m)
}
