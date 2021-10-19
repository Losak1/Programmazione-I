/*
Scrivere un programma che legge da standard input un intero n
e stampa a video se il numero è pari o dispari.
*/

package main

import "fmt"

func main() {

	var n int
	fmt.Println("Inserisci un numero: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("Il numero è pari")
	} else {
		fmt.Println("Il numero è dispari")
	}
}
