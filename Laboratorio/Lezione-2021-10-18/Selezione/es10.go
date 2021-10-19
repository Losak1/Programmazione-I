/*
Scrivere un programma che legga da standard input due numeri interi a e b
e calcoli il risultato della divisione a/b.
Se b è uguale a 0, il programma stampa Impossibile.
*/

package main

import "fmt"

func main() {

	var a, b int
	fmt.Println("Inserire due numeri")
	fmt.Scan(&a)
	fmt.Scan(&b)

	if b == 0 {
		fmt.Println("Impossibile")
	} else {
		fmt.Println("Quoziente =", float64(a)/float64(b))
	}
}
