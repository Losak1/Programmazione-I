/*
TROVA LA RADICE QUADRATA APPROSSIMATA PER ECCESSO DI X
*/

package main

import "fmt"

func main() {

	var x int
	fmt.Println("Inserisci un numero non negativo")
	fmt.Scan(&x)

	r := 0
	if x >= 0 {
		for r*r < x {
			r++
		}
	} else {
		fmt.Println("Il numero è negativo")
	}
	fmt.Print("La radice quadrata arrotondata per eccesso di ", x, " è ", r, "\n")
}
