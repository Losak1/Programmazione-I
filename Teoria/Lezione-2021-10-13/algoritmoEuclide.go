/*
ALGORTIMO DI EUCLIDE PER IL CALCOLO DEL MCD
*/

package main

import "fmt"

func main() {

	var x, y int
	fmt.Println("Inserire due numeri")
	fmt.Scan(&x)
	fmt.Scan(&y)

	for y != 0 {
		r := x % y
		x, y = y, r
	}

	fmt.Println("Il resto della divisione è", x)
}
