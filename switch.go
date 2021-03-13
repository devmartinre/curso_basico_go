package main

import "fmt"

func main() {

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es Par")
	default:
		fmt.Println("Es Impar")
	}

	// Sin condición
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value > 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condición")
	}
}
