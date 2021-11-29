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

	valid := true
	var y int

	for i := 1; i < len(os.Args); i++ {
		x, err := strconv.Atoi(os.Args[i])
		if err != nil { //ESCE SE NON INTERO
			fmt.Println("Valore in posizione", i-1, "non valido.")
			valid = false
			break
		}

		if i == 1 { //IL PRIMO ELEMENTO NON E' SOGGETTO AD ULTERIORI CONTROLLI
			y = x
			continue
		}

		if i%2 == 0 { //PER I NUMERI CHE SONO NELLA POSIZIONE REALE DISPARI (OFFSET DI 1)
			if x >= y { //DEVONO RISPETTARE LA CONDIZIONE DI ESSERE < DEL PRECEDENTE
				fmt.Println("Valore in posizione", i-1, "non valido.")
				valid = false
				break
			}
		} else { //PER I NUMERI CHE SONO NELLA POSIZIONE REALE PARI
			if x <= y { //DEVONO RISPETTARE LA CONDIZIONE DI ESSERE > DEL PRECEDENTE
				fmt.Println("Valore in posizione", i-1, "non valido.")
				valid = false
				break
			}
		}

		y = x

	}

	if valid {
		fmt.Println("Sequenza valida.")
	}
}
