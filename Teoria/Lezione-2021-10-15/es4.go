/*
CALCOLA LA SOMMA DI UNA SEQUENZA DI NUMERI TERMINANTE CON UN NUMERO NEGATIVO
*/

package main

import "fmt"

func main() {

	var x, s int

	fmt.Scan(&x)
	for x >= 0 {
		s += x
		fmt.Scan(&x)
	}
	fmt.Println("La somma è", s)
}
