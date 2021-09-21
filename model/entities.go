package model

import (
	"strconv"
)

const minLen = 5

type Result struct {
	Type   string
	Length int
	Value  string
}

func NewResult(Type string, Length int, Value string) Result {
	return Result{Type, Length, Value}
}

func GetStructure(chain string) Result /*, error*/ {
	if len(chain) >= minLen {
		Type := chain[0:2]
		LengthStr := chain[2:4]
		Value := chain[4:]
		LengthInt, err := strconv.Atoi(LengthStr) //convierte string en int
		if err == nil && LengthInt == len(Value) {
			return NewResult(Type, LengthInt, Value)
		} else {
			//return NewResult("Error", 00, "Error")
			//return err
		}
	} else {
		//return NewResult("Error", 00, "Error")
		//return err
	}

}
