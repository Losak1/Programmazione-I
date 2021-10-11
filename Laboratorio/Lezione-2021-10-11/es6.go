package main

import "fmt"

func main() {

	var base, altezza int

	fmt.Print("Inserisci la base: ")
	fmt.Scan(&base)
	fmt.Print("Inserisci la base: ")
	fmt.Scan(&altezza)

	fmt.Println("Perimetro =", 2*(base+altezza))
	fmt.Println("Area =", base*altezza)

}
