package main

import (
	. "fmt"
	//. "math"
	"math/rand"
	"time"
)

//generador de numeros al azar **copperMan**
func randInt() int {
	rand.Seed(time.Now().UTC().UnixNano())
	nrd := -1 + rand.Intn(3)
	for nrd == 0 {
		nrd = -1 + rand.Intn(3)
	}
	return nrd
}

func main() {
	//a1 := []float64{}
	//var nr float64
	a1 := []int{}
	var nr int
	r := rand.New(rand.NewSource(99))
	//r2:=rand.Int63n()
	for i := 0; i < 51; i++ {
		//nr = r.Int31n(5) //entre 0 y (n), ojo int31 = int32 para las variables
		//nr = r.Intn(5) //entre 0 y (n), ojo int = intpara las variables
		//nr = (-1 + r.Intn(3)) //genera un numero entre -1+(numeros del 0 al 2)
		//nr = (rand.Float64()*2 - 1) * 5//similar a la anterior
		//nr = rand.Float64()*(4.7 - -3.5) + -3.5
		//nr = (r.NormFloat64())*1 + 3
		//nr = Floor((rand.NormFloat64()) + 0.5)
		Println(nr)
		a1 = append(a1, nr)
	}

	cont := 0
	//var sum float64
	var sum int
	for _, v := range a1 {
		sum += v
		if v == 1 {
			cont++
		}
	}
	//Println("promedio = ", sum/float64(len(a1)))
	Println("promedio = ", sum/(len(a1)))
	Println("porcentaje de 1 : ", (cont*100)/50)

}
