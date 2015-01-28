package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = int(f)
	fmt.Println(x, y, z, f)
}
