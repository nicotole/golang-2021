package main

import (
	"entregableGo/model"
	"fmt"
)

func main() {
	p := model.NewResult("Nico", "Sape", 10)
	fmt.Println(p)
	fmt.Println(model.GetStructure("Nico"))
	fmt.Println("----------------")
	fmt.Print(model.GetStructure("NN040001"))
}
