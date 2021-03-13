package main

import "fmt"

func normalFunction(messsage string) {
	fmt.Println(messsage)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola Mundo")
	tripleArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, _ := doubleReturn(2)
	fmt.Println("value 1: ", value1)
}
