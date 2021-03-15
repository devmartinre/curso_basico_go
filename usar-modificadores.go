package main

import f "fmt"

// CarPublic Car con acceso público
type CarPublic struct {
	Brand string
	Year  int
}

// PrintMessage imprime un  mensaje
func PrintMessage(text string) {
	f.Println(text)
}

func main() {
	var myCar CarPublic
	myCar.Brand = "Ferrari"
	f.Println(myCar)

	PrintMessage("Hola platzi")
}
