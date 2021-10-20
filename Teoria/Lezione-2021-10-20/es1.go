/*
STAMPA GLI ELEMENTI DELLA SEQUENZA DI COLLATZ PARTENDO DA UN NUMERO INSERITO IN INPUT
*/

package main

import "fmt"

func collatz(x int) int {
	if x%2 == 0 {
		return x / 2
	} else {
		return 3*x + 1
	}
}

func main() {
	var x int
	fmt.Scan(&x)
	i := 0
	for ; x != 1; i++ {
		x = collatz(x)
		fmt.Println(i, ":", x)
	}
}
