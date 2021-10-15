/*
STAMPA LE ULTIME 3 CIFRE DI UN NUMERO
*/

package main

import "fmt"

func main() {

	var x int
	fmt.Println("Inserisci un numero maggiore o uguale a 1000")
	fmt.Scan(&x)

	if x >= 1000 {
		if r := x % 1000; r == 0 {
			fmt.Println("Le ultime tre cifre sono zeri")
		} else {
			if r < 10 {
				fmt.Print("Le ultime tre cifre sono 00", r, "\n")
			} else if r < 100 {
				fmt.Print("Le ultime tre cifre sono 0", r, "\n")
			} else {
				fmt.Print("Le ultime tre cifre sono ", r, "\n")
			}
		}
	} else {
		fmt.Println("Non è maggiore o uguale a 1000")
	}
}
