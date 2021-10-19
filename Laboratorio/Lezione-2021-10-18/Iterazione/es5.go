/*
Scrivere un programma che legga da standard input due numeri interi a e b
e stampi a video la somma dei numeri pari compresi tra a e b (estremi esclusi).
*/

package main

import "fmt"

func main() {

	var a, b int
	fmt.Println("Inserisci due numeri: ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a > b {
		a, b = b, a
	}
	if a%2 == 0 {
		/*nel caso sia pari, lo trasformo nel dispari successivo,
		considerando come è stato strutturato l'ingresso nel ciclo*/
		a++
	}
	somma := 0
	for i := a + 1; i < b; i += 2 {
		somma += i
	}
	fmt.Println("Somma =", somma)
}
