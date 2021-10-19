/*
Scrivere un programma che legge da standard input un numero intero n
(specificato senza segno se maggiore o uguale a 0)
e stampi a video il numero con segno.
*/

package main

import "fmt"

func main() {

	var n int
	fmt.Println("Inserisci un numero: ")
	fmt.Scan(&n)

	if n > 0 {
		fmt.Print("+", n, "\n")
	} else {
		fmt.Println(n)
	}
}
