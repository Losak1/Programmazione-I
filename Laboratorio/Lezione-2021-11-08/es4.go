/*
TESTO DEL PROBLEMA
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
)

const EPSILON = 1.0e-9

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è maggiore
di `y` di almeno una costante EPSILON */
/* ÈXMaggioreDiY(5.0,4.999) -> true */
/* ÈXMaggioreDiY(5.0,4.9999999999) -> false */
func ÈXMaggioreDiY(x, y float64) bool {
	return (x - y) > EPSILON
}

/* La funzione `func ÈXUgualeAY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è uguale
a `y` a meno di una costante EPSILON */
/* ÈXUgualeAY(5.0,4.999) -> false */
/* ÈXUgualeAY(5.0,4.9999999999) -> true */
func ÈXUgualeAY(x, y float64) bool {
	return math.Abs(x-y) <= EPSILON
}

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è minore
di `y` di almeno una costante EPSILON */
/* ÈXMinoreDiY(4.999,5.0) -> true */
/* ÈXMinoreDiY(4.9999999999,5.0) -> false */
func ÈXMinoreDiY(x, y float64) bool {
	return (x - y) < -EPSILON
}

func main() {
	var s int64
	fmt.Println("Inserisci un seme per la generazione casuale:")
	fmt.Scan(&s)
	rand.Seed(s)

	var m, q float64
	fmt.Println("Inserisci m e q per la retta:")
	fmt.Scan(&m, &q)

	for i := 0; i < 10; i++ {
		px := rand.Float64() * 10
		py := rand.Float64() * 10

		lineY := m*px + q

		fmt.Print("Punto (", px, ",", py, ") - ")
		if ÈXMaggioreDiY(py, lineY) {
			fmt.Print("Il punto sta sopra la retta\n")
		} else if ÈXUgualeAY(py, lineY) {
			fmt.Print("Il punto è sulla la retta\n")
		} else {
			fmt.Print("Il punto sta sotto la retta\n")
		}

	}
}
