package main

import (
	"entregableGo/model"
	"fmt"
)

func main() {
	fmt.Println(model.GetStructure("Nico"))
	fmt.Println("----------------")
	fmt.Print(model.GetStructure("NN040001"))
	fmt.Println("----------------")
	fmt.Print(model.GetStructure("TX03ABC"))
	fmt.Println("----------------")
	fmt.Print(model.GetStructure("TX02AB7"))
	fmt.Println("----------------")
	fmt.Print(model.GetStructure("TX0JABC"))
}
