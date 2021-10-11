package main

import (
	"fmt"
	"math"
)

func main() {

	var eta1, eta2 int

	fmt.Print("Età persona 1 = ")
	fmt.Scan(&eta1)
	fmt.Print("Età persona 2 = ")
	fmt.Scan(&eta2)

	var sommaEta, mediaEta, mediaArrotondataDifetto, mediaArrotondataEccesso float64

	sommaEta = float64(eta1 + eta2)
	fmt.Println("Somma età = ", sommaEta)

	mediaEta = float64(eta1+eta2) / 2
	fmt.Println("Media età = ", mediaEta)

	mediaArrotondataDifetto = math.Floor(mediaEta)
	fmt.Println("Media età arrotondata per difetto all'intero inferiore = ", mediaArrotondataDifetto)

	mediaArrotondataEccesso = math.Ceil(mediaEta)
	fmt.Println("Media età arrotondata per eccesso all'intero superiore = ", mediaArrotondataEccesso)

	fmt.Println("Somma età a 10 anni = ", sommaEta+20)
	fmt.Println("Media età a 10 anni = ", mediaEta+10)

}
