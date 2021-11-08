/*
ESEMPI SU SWITCH
*/

package main

import "fmt"

func main() {

	var x, r int
	fmt.Scan(&x)
	fmt.Scan(&r)

	switch x {
	case r * r: //caso d'uso di un'espressione numerica (poco usata)
		fmt.Println("Hai inserito la radice corretta")
	default:
		fmt.Println("Peccato")
	}
}
