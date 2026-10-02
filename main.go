package main

import (
	"fmt"
	ecc "github.com/zzzyoonnn/coinflow/elliptic-curve"
	"math/big"
	//"math/rand" use 'Finite field'
)

/* Finite field
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

// p is field order, p = 7, 11, 17, 19, 31
// for every non-zero element k, compute k ^ (p - 1) mod p equals 1
// => {1 ^ (p - 1) % p, 2 ^ (p - 1) % p, ..., (p - 1) ^ (p - 1) % p}

// for any element k in the field with order => k ^ (p - 1) % p == 1

func ComputeFieldOrderPower() {
	orders := []int{7, 11, 17, 19, 31}
	for _, p := range orders {
		fmt.Printf("value of p is %d\n", p)
		for i := 1; i < p; i++ {
			elm := ecc.NewFieldElement(big.NewInt(int64(p)), big.NewInt(int64(i)))
			fmt.Printf("for element %vits power of p - 1 is %v\n", elm, elm.Power(big.NewInt(int64(p-1))))
		}
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

	ComputeFieldOrderPower()

	f2 := ecc.NewFieldElement(big.NewInt(int64(19)), big.NewInt(int64(2)))
	f7 := ecc.NewFieldElement(big.NewInt(int64(19)), big.NewInt(int64(7)))
	fmt.Printf("field element 2 / 7 with order 19 is %v", f2.Divide(f7))

	f46 = ecc.NewFieldElement(big.NewInt(int64(59)), big.NewInt((int64(46))))
	fmt.Printf("field element 46 * 46 with order 59 is %v", f46.Multiply(f46))
	fmt.Printf("field element 46 with power of 60 is %v", f46.Power(big.NewInt(int64(60))))
	// 60 % (59 - 1) = 2
}
*/

// Elliptic curve
func main() {
	x1 := ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(192)))
	y1 := ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(105)))
	a := ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(0)))
	b := ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(7)))
	p1 := ecc.NewEllipticCurvePoint(x1, y1, a, b)
	fmt.Printf("elliptic curve point over finite field is %s\n", p1)

	yNeg := y1.Negate()
	p2 := ecc.NewEllipticCurvePoint(x1, yNeg, a, b)
	res := p1.Add(p2)
	fmt.Printf("addition of points on vertical line over finite field is %s\n", res)

	x2 := ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(17)))
	y2 := ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(56)))
	p2 := ecc.NewEllipticCurvePoint(x2, y2, a, b)
	res = p1.Add(p2)
	fmt.Printf("p1 + p2 over field order 223 is %s\n", res)

	res = p1.Add(p1)
	fmt.Printf("p1 + p1 over field order 223 is %s\n", res)

	// (170, 142) + (60, 139)
	x1 = ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(170)))
	y1 = ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(142)))
	p1 = ecc.NewEllipticCurvePoint(x1, y1, a, b)
	x2 = ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(60)))
	y2 = ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(139)))
	p2 = ecc.NewEllipticCurvePoint(x2, y2, a, b)
	res = p1.Add(p2)
	fmt.Printf("p1(170, 142) + p2(60, 139) over field order 223 is %s\n", res)

	// (47, 71) + (17, 56)
	x1 = ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(47)))
	y1 = ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(71)))
	p1 = ecc.NewEllipticCurvePoint(x1, y1, a, b)
	x2 = ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(17)))
	y2 = ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(56)))
	p2 = ecc.NewEllipticCurvePoint(x2, y2, a, b)
	res = p1.Add(p2)
	fmt.Printf("p1(47, 71) + p2(17, 56) over field order 223 is %s\n", res)

	// (143, 98) + (76, 66)
	x1 = ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(143)))
	y1 = ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(98)))
	p1 = ecc.NewEllipticCurvePoint(x1, y1, a, b)
	x2 = ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(76)))
	y2 = ecc.NewFieldElement(big.NewInt(int64(223)), big.NewInt(int64(66)))
	p2 = ecc.NewEllipticCurvePoint(x2, y2, a, b)
	res = p1.Add(p2)
	fmt.Printf("p1(143, 98) + p2(76, 66) over field order 223 is %s\n", res)
}
