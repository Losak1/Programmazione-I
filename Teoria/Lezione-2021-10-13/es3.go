/*
VERIFICA CHE LA SOMMA DELLE ULTIME 3 CIFRE DI UN NUMERO, SIANO MAGGIORI DI UNA SOGLIA S
*/

package main

import "fmt"

func main() {

	var x, s int
	fmt.Println("Inserisci un numero maggiore o uguale a 1000 e una soglia")
	fmt.Scan(&x)
	fmt.Scan(&s)

	if x >= 1000 {
		if t := x%10 + (x/10)%10 + (x/100)%10; t > s {
			fmt.Println("La somma delle prime 3 cifre è maggiore della soglia")
		} else {
			fmt.Println("La somma delle prime 3 cifre non è maggiore della soglia")
		}
	} else {
		fmt.Println("Non è maggiore o uguale a 1000")
	}
}
