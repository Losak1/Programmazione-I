/*
Scrivere un programma che legga da standard input 4 valori a virgola mobile:

i primi due valori sono il coefficiente angolare m e
il termine noto q di una retta r: y = m*x + q
il terzo e il quarto valore sono le coordinate px e py di un punto P(px,py)
Il programma deve determinare se il punto P sta sopra o sotto la retta
od appartiene ad essa, e stampare a video il relativo messaggio.
*/

package main

import "fmt"

func main() {

	var m, q, px, py float64
	fmt.Print("Inserisci m e q: ")
	fmt.Scan(&m)
	fmt.Scan(&q)
	fmt.Print("Inserisci x e y: ")
	fmt.Scan(&px)
	fmt.Scan(&py)

	sly := m*px + q

	if py > sly {
		fmt.Println("Il punto sta sopra la retta")
	} else if py == sly {
		fmt.Print("Il punto appartiene alla retta")
	} else {
		fmt.Print("Il punto sta sotto la retta")
	}
}
