/*
CONTA CON QUANTI ZERI CI SONO NEL NUMERO INSERITO
*/

package main

import "fmt"

func main() {

	var x int
	fmt.Println("Inserisci un numero maggiore di zero:")
	fmt.Scan(&x)
	t := x

	if x > 0 {
		i := 0
		for x != 0 {
			if x%10 == 0 {
				i++
			}
			x /= 10
		}
		fmt.Println(t, "ha", i, "zeri")
	} else {
		fmt.Println(t, "non è maggiore di zero")
	}
}
