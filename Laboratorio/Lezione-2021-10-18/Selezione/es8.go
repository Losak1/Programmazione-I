/*
Scrivere un programma che legge da standard input un numero intero
e stampa "Fizz" se il numero è multiplo di 3, "Buzz" se il numero
è multiplo di 5, "Fizz Buzz" se è multiplo sia di 3 sia di 5, niente altrimenti.
*/

package main

import "fmt"

func main() {

	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	/*
		//LASCIA UNO SPAZIO EXTRA ACCANTO A FIZZ
		if n%3 == 0 {
			fmt.Print("Fizz ")
		}

		if n%5 == 0 {
			fmt.Print("Buzz")
		}*/
	if n%3 == 0 && n%5 == 0 { //n%15 == 0
		fmt.Print("Fizz Buzz")
	} else if n%3 == 0 {
		fmt.Print("Fizz")
	} else if n%5 == 0 {
		fmt.Print("Buzz")
	}
}
