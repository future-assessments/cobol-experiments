package main

import "fmt"

func main() {
	fmt.Println("Hello, compute")
	fpl_pi := 3.14159265358979323
	fpl_rad := 2.0
	fpl_cir := fpl_pi * (2 * fpl_rad)
	message := fmt.Sprintf("Perimeter %10.15f   Radius %2.16f Pi: %3.18f\n", fpl_cir, fpl_rad, fpl_pi)
	fmt.Println(message)
}

