/*
TESTO DEL PROBLEMA
*/

package main

import "fmt"

func main() {
	var a, b float64
	var segno string

	fmt.Println("inserisci un'operazione aritmetica")
	fmt.Scan(&a, &segno, &b)

	var risultato float64

	switch segno {
	case "+":
		risultato = a + b
		fmt.Println("Risultato:", risultato)
	case "-":
		risultato = a - b
		fmt.Println("Risultato:", risultato)
	case "*":
		risultato = a * b
		fmt.Println("Risultato:", risultato)
	case "/":
		risultato = a / b
		fmt.Println("Risultato:", risultato)
	default:
		fmt.Println("Operatore non riconosciuto")
	}
}
