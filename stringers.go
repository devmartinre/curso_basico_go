package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (mypc pc) String() string {
	return fmt.Sprintf("Texto %d GB RAM, %d GB Disco y es una %s", mypc.ram, mypc.disk, mypc.brand)
}

func main() {
	mypc := pc{ram: 16, brand: "msi", disk: 100}
	fmt.Println(mypc)
}
