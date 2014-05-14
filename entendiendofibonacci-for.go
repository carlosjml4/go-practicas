//entra el numero binario
package main

import "fmt"

func main() {
	x, y := 0, 1
	for y < 100 {
		fmt.Println(x)
		x, y = y, x+y
	}
}
