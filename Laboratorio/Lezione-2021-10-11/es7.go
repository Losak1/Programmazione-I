package main

import (
	"fmt"
	"math"
)

func main() {

	var raggio float64

	fmt.Print("Raggio = ")
	fmt.Scan(&raggio)

	fmt.Println("Circonferenza =", 2*raggio*math.Pi)
	fmt.Println("Area =", raggio*raggio*math.Pi)

}
