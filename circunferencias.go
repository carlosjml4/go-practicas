package main

import (
	"fmt"
	"math"
)

type circ struct {
	x, y, r float64
	cs      string
}

func RelaCir(x1, y1, r1 float64, c1 string, x2, y2, r2 float64, c2 string) {
	var d float64
	var st string
	d = math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
	if d == 0 && (r1 > r2 || r2 > r1) {
		st = "es Interior Concentrica"
	} else if math.Floor(d*10+0.5)/10 == math.Floor((r1+r2)*10+0.5)/10 {
		st = "es Tangente Exterior"
	} else if math.Floor(d*10+0.5)/10 == math.Floor((r1-r2)*10+0.5)/10 {
		st = "es Tangente Interior"
	} else if d > 0 && d < (r1-r2) {
		st = "es Interior"
	} else if d > (r1 + r2) {
		st = "es Exterior"
	} else {
		st = "es secante"
	}
	fmt.Println("La relacion de ", c1, "y", c2, st)
}

func main() {
	acirc := []circ{}
	var (
		x, y, r float64
		nc      int
		elc     circ
	)

	fmt.Println(" Ingrese datos (x,y,r) para las circunferencias", "\n",
		"separe los datos con espacio y finalice con enter,", "\n",
		"los numeros deben ser decimales o enteros mayores a 0.", "\n")

	fmt.Print("Numero de circunferencias?: ", "\n")
	fmt.Scan(&nc)

	for i := 0; i < nc; i++ {
		fmt.Print("Circunferencia ", string(i+65), ":")
		_, er := fmt.Scan(&x, &y, &r)
		if er != nil {
			fmt.Println("Datos incorrectos intente nuevamente")
			main()
		} else {
			elc.x = x
			elc.y = y
			elc.r = r
			elc.cs = string(i + 65)
			acirc = append(acirc, elc)
		}
	}

	for i := 0; i < nc; i++ {
		for j := i + 1; j < nc; j++ {
			RelaCir(acirc[i].x, acirc[i].y, acirc[i].r, acirc[i].cs,
				acirc[j].x, acirc[j].y, acirc[j].r, acirc[j].cs)
		}
	}
}
