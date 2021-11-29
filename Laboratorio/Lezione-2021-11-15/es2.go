/*
TESTO DEL PROBLEMA
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	prodotto := 1
	for _, v := range os.Args {
		x, err := strconv.Atoi(v)
		if err == nil { //VALIDO SE INTERO
			prodotto *= x
		}
	}

	fmt.Println("Il risultato della moltiplicazione tra i numeri interi è", prodotto)
}
