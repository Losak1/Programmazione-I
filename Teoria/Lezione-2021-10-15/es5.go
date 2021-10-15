/*
TESTO DEL PROBLEMA
*/

package main

import "fmt"

func main() {

	var n int
	fmt.Println("Inserisci il numero da indovinare: ")
	fmt.Scan(&n)

	var t int
	fmt.Print("Inserisci tentativo")
	fmt.Scan(&t)

	for t != n {
		if t < n {
			fmt.Println("Troppo piccolo, ritenta")
		} else {
			fmt.Println("Troppo grande, ritenta")
		}
		fmt.Scan(&t)
	}
	fmt.Println("Hai indovinato!")
}
