package main

import "fmt"

func main() {

	var n1, n2, n3, n4 int

	fmt.Scan(&n1, &n2, &n3, &n4)

	fmt.Println(n1, n2, n3, n4)
}

/*
-Cosa succede se inserite l'input su righe diverse?

Il software si comporta correttamente.

-Cosa succede se inserite meno numeri di quelli richiesti?

Il software rimane in attesa del quarto elemento, senza terminare
la sua esecuzione. In caso di chiusura dello standard input (CTRL+D)
l'ultima variabile assume valore 0.

-Cosa succede se inserite più numeri di quelli richiesti?

Il programma prende in input solo i primi 4,
lasciando quelli in eccesso nel buffer.

-Cosa succede se inserite un valore diverso da un numero intero
(numero reale, lettera, parola, ...)

[...]

*/
