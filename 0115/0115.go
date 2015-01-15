package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("The time is ", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.", math.Nextafter(2,3))
	fmt.Println(math.Pi)
	fmt.Println(math.Pi * 2)
}



