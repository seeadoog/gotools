package goutils

import (
	"fmt"
	"testing"
)

func TestSetValueByExp(t *testing.T) {
	o := make(map[string]interface{})
	SetValueByExp(o,"a=5")
	SetValueByExp(o,"b=6")
	SetValueByExp(o,"c.d=hh")
	SetValueByExp(o,"e.f.h.k=ll")

	fmt.Println(o)
}

