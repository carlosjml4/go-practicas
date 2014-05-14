package main

import (
	"fmt"
)
m := image.NewRGBA(image.Rect(0, 0, 640, 480))
blue := color.RGBA{0, 0, 255, 255}
draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)




func main() {
	//filas y columnas ==
	fl := 0
	fmt.Println("Introduce lado de la array")
	fmt.Scan(&fl)
	//definimos array 2d
	a2 := make([][]int, fl)
	for i := range a2 {
		a2[i] = make([]int, fl)
	}
	LlenaArrayCaracol(a2, fl)
	ImpArry(a2, fl)
}
