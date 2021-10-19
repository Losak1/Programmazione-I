/*
Scrivere un programma che legge da standard input
un numero intero n e verifica se il numero è multiplo di 10.
*/

package main

import "fmt"

func main() {

	var n int
	fmt.Println("Inserisci un numero: ")
	fmt.Scan(&n)

	if n%10 == 0 {
		fmt.Println("Il numero è multiplo di 10")
	} else {
		fmt.Println("Il numero non è multiplo di 10")
	}
}
