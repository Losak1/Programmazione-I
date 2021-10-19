/*
Scrivere un programma che legga da standard input le ampiezze di due angoli
di un triangolo e stampi, se possibile, l'ampiezza del terzo angolo.
*/

package main

import "fmt"

func main() {

	var a, b int
	fmt.Print("Inserire le ampiezze dei due angoli: ")
	fmt.Scan(&a)
	fmt.Scan(&b)

	if a+b >= 180 {
		fmt.Println("I due angoli non appartengono ad un triangolo")
	} else {
		fmt.Print("Ampiezza terzo angolo = ", 180-a-b, "°\n")
	}
}
