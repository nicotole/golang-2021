package model

import (
	"fmt"
	"testing"
)

func TestInputChains(t *testing.T) {
	//chainTest := "TX040001"
	/*chainTest := "TB020"
	result, _ := model.GetStructure(chainTest)
	st := model.NewResult("TX", 04, "0001")
	if result != st {
		fmt.Println("No funciona")
	} */

	chainTest := "TX040001"
	//chainTest := "TB020"
	result, err := GetStructure(chainTest)
	fmt.Println("JP recomendation")
	if result == nil {
		t.Error(err)
	} else {
		t.Log(result)
	}

}
