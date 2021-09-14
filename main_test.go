package main

import (
	"entregableGo/model"
	"fmt"
	"testing"
)

func TestInputChains(t *testing.T) {
	chainTest := "TX040001"
	result := model.GetStructure(chainTest)
	st := model.NewResult("TX", 04, "0001")
	if result != st {
		fmt.Println("No funciona")
	}
}
