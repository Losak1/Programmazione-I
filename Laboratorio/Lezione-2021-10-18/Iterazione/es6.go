/*
Scrivere un programma che, dopo aver letto da standard input un numero intero n,
chiede all'utente di inserire n numeri interi (sempre da standard input).

Dopo aver letto gli n numeri interi, il programma deve stampare:

la somma degli n numeri letti;
il minimo tra i numeri letti;
il massimo tra i numeri letti;
il numero di interi letti strettamente positivi (maggiori di 0),
strettamente negativi (minori di 0), e nulli.
*/

package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	fmt.Println("Inserisci", n, "numeri: ")
	var x, minimo, massimo int
	somma := 0
	nPositivi := 0
	nNegativi := 0
	nNulli := 0

	fmt.Scan(&x)
	somma += x
	if x > 0 {
		nPositivi++
	} else if x == 0 {
		nNulli++
	} else {
		nNegativi++
	}
	minimo = x
	massimo = x

	for i := 0; i < n-1; i++ {
		fmt.Scan(&x)
		somma += x
		if x > 0 {
			nPositivi++
		} else if x == 0 {
			nNulli++
		} else {
			nNegativi++
		}
		if x < minimo {
			minimo = x
		} else if x > massimo {
			massimo = x
		}
	}

	fmt.Println("Somma =", somma)
	fmt.Println("Valore minimo =", minimo)
	fmt.Println("Valore massimo =", massimo)
	fmt.Println("Interi > 0 =", nPositivi)
	fmt.Println("Interi < 0 =", nNegativi)
	fmt.Println("Interi = 0 =", nNulli)
}
