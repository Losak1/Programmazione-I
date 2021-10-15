/*
STAMPA I NUMERI DA 0 A N-1
*/
package main

import "fmt"

func main() {

	var n int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}
