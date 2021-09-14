package main

import (
	"entregableGo/model"
	"fmt"
)

func main() {
	p := model.NewResult("Nico", 10, "Sape")
	fmt.Println(p)
	fmt.Println(model.GetStructure("Nico"))
	fmt.Println("----------------")
	fmt.Print(model.GetStructure("NN040001"))
	fmt.Println("----------------")
	fmt.Print(model.GetStructure("TX03ABC"))
	fmt.Print(model.GetStructure("TX03ABC"))
}
