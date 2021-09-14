package main

import (
	"entregableGo/model"
	"fmt"
)

func main() {
	p := model.NewResult("Nico", "Sape", 10)
	fmt.Println(p)
}
