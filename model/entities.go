package model

import (
	"strconv"
)

type Result struct {
	Type   string
	Value  string
	Length int
}

func NewResult(Type string, Value string, Length int) Result {
	return Result{Type, Value, Length}
}

func GetStructure(chain string) Result {
	Type := chain[0:2]
	LengthStr := chain[2:4]
	Value := chain[4:]
	LengthInt, err := strconv.Atoi(LengthStr) //convierte string en int
	if err == nil {
		return NewResult(Type, Value, LengthInt)
	} else {
		return NewResult("Error", "Error", 00)
	}

}
