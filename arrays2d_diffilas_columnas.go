ajuego := make([][]int, 5)
for i := range ajuego {
	ajuego[i] = make([]int, 3)
}
//llenamos con numeros al azar
for i := 0; i < 5; i++ {
	for j := 0; j < 3; j++ {
		ajuego[i][j] = rand.Intn(10)
	}
}
for i := 0; i < 5; i++ {
	Print("[")
	for j := 0; j < 3; j++ {
		Printf("%2d", ajuego[i][j])
	}
	Println("]")
}
