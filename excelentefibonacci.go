//entra el numero binario
package main

import "fmt"

func main() {
	//for i, j := 0, 1; j < 20; i, j = i+j, i {
	//	fmt.Println(i)
	//}
	//for i := 0; i < 1; {
	i := 0
	for j := 1; j < 20; {
		fmt.Println("i", i)
		fmt.Println("j", j)
		i = i + j
		j = i + j
	}

}
