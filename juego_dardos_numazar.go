package main

import (
	. "fmt"
	"math/rand"
	"strconv"
	"time"
)

type jugadas struct {
	lanz1, lanz2, lanz3, tlanz int
}

//generador de numeros al azar **copperMan**
func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

//imprime tabla resultado
func tblResult(a [][]jugadas, njs, j, mp int) {
	//titulos de columnas
	Print("             ")
	for i := 0; i < njs; i++ {
		Print(" JUGADOR  " + strconv.Itoa(i+1) + "   ")
	}
	Println()
	//despliega puntaje
	for i := 0; i < 20; i++ {
		Printf("jugada: %2d [", i+1)
		for j := 0; j < njs; j++ {
			Printf("%2d", a[i][j])
		}
		Println("]")
	}
	Println()
	Println("EL GANADOR ES!!!: ", "El Jugador No. ", j, "con", mp, "puntos.")
}

//realiza el juego
func jugar(ad []int, ap [][]jugadas, njs int) (int, int) {
	var jgda jugadas
	lim := 501 //limite juego
	lz1, lz2, lz3 := 0, 0, 0
	//crea variables con el limite de puntage para c/a jugador
	//para luego ir disminuyendo con los lanzamientos
	aux := make([]int, njs)
	for i := 0; i < njs; i++ {
		aux[i] = lim
	}
	//funcion que verifica el lanzamiento con el puntaje
	//devolviendo el puntaje de lanzamiento
	v_lanz := func(vl_aux int) (int, int, int) {
		lz1 = ad[randInt(0, len(ad)-1)] //numeros al azar
		lz2 = ad[randInt(0, len(ad)-1)] //para cada uno
		lz3 = ad[randInt(0, len(ad)-1)] //de los lanzamientos
		//verifica y descuenta del limite el puntaje optenido
		if vl_aux-lz1 == 0 {
			lz2, lz3 = 0, 0
		} else if vl_aux-lz1 < 0 {
			lz1, lz2, lz3 = 0, 0, 0
		} else if (vl_aux-lz1)-lz2 == 0 {
			lz3 = 0
		} else if (vl_aux-lz1)-lz2 < 0 {
			lz1, lz2, lz3 = 0, 0, 0
		} else if (vl_aux-lz1-lz2)-lz3 == 0 {
			return lz1, lz2, lz3
		} else if (vl_aux-lz1-lz2)-lz3 < 0 {
			lz1, lz2, lz3 = 0, 0, 0
		}
		return lz1, lz2, lz3
	}
	//ciclo de lanzamientos
	for i := 0; i < 20; i++ {
		for j := 0; j < njs; j++ {
			jgda.lanz1, jgda.lanz2, jgda.lanz3 = v_lanz(aux[j])
			jgda.tlanz = aux[j] - (jgda.lanz1 + jgda.lanz2 + jgda.lanz3)
			aux[j] = jgda.tlanz
			ap[i][j] = jgda
		}
	}
	//devuelviendo el ganador
	men, jgdr := aux[0], 1
	for i := 1; i < njs; i++ {
		if aux[i] < men {
			men = aux[i]
			jgdr = i + 1
		}
	}
	return jgdr, men
}

func main() {
	var njs int
	ndiana := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
		18, 19, 20, 21, 22, 24, 25, 26, 27, 28, 30, 32, 33, 34, 36, 38, 39, 40,
		42, 45, 48, 50, 51, 54, 57, 60}
	Print("Ingrese numero de jugadores: ")
	Scan(&njs)
	Println()
	//se crea array para los puntajes
	aptaje := make([][]jugadas, 20)
	for i := range aptaje {
		aptaje[i] = make([]jugadas, njs)
	}

	jugador, men_ptaje := jugar(ndiana, aptaje, njs)
	tblResult(aptaje, njs, jugador, men_ptaje)
}
