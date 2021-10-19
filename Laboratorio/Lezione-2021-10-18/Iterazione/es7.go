/*
Scrivere un programma che legga da standard input una sequenza di numeri reali
strettamente positivi (un numero è strettamente positivo se è maggiore di 0;
se un numero è minore o uguale 0 non è strettamente positivo).
La lettura termina quando viene letto un numero minore o uguale a 0.

Il programma deve stampare a video il risultato della media aritmetica dei valori inseriti.
*/

package main

import "fmt"

func main() {

	var x float64
	fmt.Print("Inserisci una sequenza di numeri (interrompi con numero<=0): ")
	somma := 0.0
	n := 0

	fmt.Scan(&x)
	for x > 0 {
		somma += x
		n++
		fmt.Scan(&x)
	}

	fmt.Println("Media aritmetica:", somma/float64(n))
}
