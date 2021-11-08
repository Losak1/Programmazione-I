/*
ESEMPI SU SWITCH
*/

package main

import "fmt"

func main() {

	var x int
	fmt.Scan(&x)

	switch x {
	case 0, 1, 2:
		fmt.Println("Minore di 3")
	case 4:
		fmt.Println("È quattro")
		fallthrough
		/*
			su go non ci sono break nello switch
			si esce automaticamente dal primo blocco valido
			è possibile di continuare l'esecuzione evitando l'uscita
			attraverso il comando fallthrough
		*/
	default:
		fmt.Println("Non ho informazioni a riguardo")
	}
}
