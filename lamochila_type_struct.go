package main

import (
	"fmt"
)

//var ctr int

func a(al []string) {
	var Atmp = []string{}
	var Atmp0 = []string{}
	var i, cont int
	var h func()
	ne := len(al)
	var s string

	h = func() {

		for ; i < ne; i++ {
			Atmp = append(Atmp, s+al[i])
			Atmp0 = append(Atmp0, s+al[i])
		}
	}

	for j := 0; j < 15; j++ {

		h()
		s = Atmp[j]

		//if cont < ne {
		if j < ne {
			i = cont + 1

		} else if j%ne <= ne {
			cont = j / ne
			i = cont + 3
		}
		cont++
	}

	for i := 0; i < ne; i++ {
		s = al[i]
		Atemp = append(Atemp, s)
	}
	for i := 0; i < ne-1; i++ {
		for j := i + 1; j < ne; j++ {
			s = al[i] + al[j]
			Atemp = append(Atemp, s)
		}
	}

	for i := 0; i < ne-1; i++ {
		for j := i + 1; j < ne; j++ {
			for k := j + 1; k < ne; k++ {
				s = al[i] + al[j] + al[k]
				Atemp = append(Atemp, s)
			}
		}
	}

	for i := 0; i < ne-1; i++ {
		for j := i + 1; j < ne; j++ {
			for k := j + 1; k < ne; k++ {
				for l := k + 1; l < ne; l++ {
					s = al[i] + al[j] + al[k] + al[l]
					Atemp = append(Atemp, s)
				}
			}
		}
	}

	fmt.Println(Atmp, len(Atmp))
}

func main() {
	al := make([]string, 5)
	al = []string{"A", "B", "C", "D", "E"}

	a(al)

}
