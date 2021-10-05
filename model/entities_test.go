package model

import (
	"fmt"
	"testing"
)

func TestInputChains(t *testing.T) {
	chainTest := "TX040001"
	result, err := GetStructure(chainTest)
	fmt.Println("JP recomendation")
	if result == nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
