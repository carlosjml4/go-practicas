package main

import (
	"fmt"
	"math"
)

func main() {
	var e float64

	e = 1.5437
	fmt.Println(math.Floor(e*1000+0.5) / 1000)
}
