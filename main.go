package main

import (
	"fmt"
	ecc "github.com/zzzyoonnn/coinflow/elliptic-curve"
)

func main() {
	f44 := ecc.NewFieldElement(57, 44)
	f33 := ecc.NewFieldElement(57, 33)
	res := f44.Add(f33)
	fmt.Printf("Field element 44 add to field element 33 is %v", res)
	// -44 negate of 44 is 57 - 44 is 33
	fmt.Printf("negete of field element 44 is %v", res.Negate())
}