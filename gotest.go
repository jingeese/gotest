package gotest

import (
	"fmt"
)

func createPointer() *float64 {
	var myFloat float64 = 34.5
	fmt.Println(myFloat)
	return &myFloat
}
