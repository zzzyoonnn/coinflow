package main

import (
	"fmt"
	ecc "github.com/zzzyoonnn/coinflow/elliptic-curve"
	"math/big"
	"math/rand"
)

func SolveField19MultiplySet() {
	// randomly select a num from 1 to 18
	min := 1
	max := 18
	k := rand.Intn(max-min+1) + min
	fmt.Printf("randomly select k is %d\n", k)
	element := ecc.NewFieldElement(big.NewInt(19), big.NewInt(int64(k)))
	for i := 0; i < 19; i++ {
		fmt.Printf("element %d multiply with %d is %v", k, i, element.ScalarMul(big.NewInt(int64(i))))
	}
}

func main() {
	f44 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(44))
	f33 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(33))
	res := f44.Add(f33)
	fmt.Printf("Field element 44 add to field element 33 is %v", res)
	// -44 negate of 44 is 57 - 44 is 33
	fmt.Printf("negete of field element 44 is %v", res.Negate())

	fmt.Printf("Field element 44 - 33 is %v", f44.Subtract(f33))
	fmt.Printf("Field element 33 - 44 is %v", f33.Subtract(f44))

	// check (46 + 44) % order == 33
	fmt.Printf("check 46 + 44 over modular of 57 is %d\n", (46+44)%57)
	// f33 - f44 = f46 => f46 + f44 == f33
	f46 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(46))
	fmt.Printf("field element 46 + 44 is %v", f46.Add(f44))

	fmt.Printf("multiply 46 with itself is %v", f46.Multiply(f46))
	fmt.Printf("element 46 with the power of 2 is %v\n", f46.Power(big.NewInt(2)))

	// {0, 1, ..., 18}, select any element from the field with order 19, compute following:
	// // {k.0, k.1, k.2, ..., k.18}, k != 0
	SolveField19MultiplySet()
}
