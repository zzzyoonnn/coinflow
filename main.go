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

	fmt.Printf("Field element 44 - 33 is %v", f44.Subtract(f33))
	fmt.Printf("Field element 33 - 44 is %v", f33.Subtract(f44))

	// check (46 + 44) % order == 33
	fmt.Printf("check 46 + 44 over modular of 57 is %d\n", (46+44)%57)
	// f33 - f44 = f46 => f46 + f44 == f33
	f46 := ecc.NewFieldElement(57, 46)
	fmt.Printf("field element 46 + 44 is %v", f46.Add(f44))

	fmt.Printf("multiply 46 with itself is %v", f46.Multiply(f46))
	fmt.Printf("element 46 with the power of 2 is %v", f46.Power(2))
}