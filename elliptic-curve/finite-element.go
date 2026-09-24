package elliptic_curve

import (
	"fmt"
	"math"
)

type FieldElement struct {
	order uint64 // field order
	num   uint64 // value of the given element in the field
}

func NewFieldElement(order uint64, num uint64) *FieldElement {
	/*
		init function for FieldElement
	*/

	if num >= order {
		err := fmt.Sprintf("Num not in the range of 0 to %d", order-1)
		panic(err)
	}

	return &FieldElement{
		order: order,
		num:   num,
	}
}

func (f *FieldElement) String() string {
	// __repr for python
	return fmt.Sprintf("FieldElement{order: %d, num: %d}\n", f.order, f.num)
}

func (f *FieldElement) EqualTo(other *FieldElement) bool {
	return f.order == other.order && f.num == other.num
}

func (f *FieldElement) checkOrder(other *FieldElement) {
	if f.order != other.order {
		panic("field element with the same order")
	}
}

func (f *FieldElement) Add(other *FieldElement) *FieldElement {
	f.checkOrder(other)

	// remember the modular
	// operator overloading for +, __add__ python
	return NewFieldElement(f.order, (f.num+other.num)%f.order)
}

/*
a, b (a + b) % order = 0, b is called negate of a, b = -a
*/
func (f *FieldElement) Negate() *FieldElement {
	/*
		b, (a + b) % order = 0, b = order - a, (a + b) % order => (a + order - a) => order => order % order = 0
	*/

	return NewFieldElement(f.order, (f.order-f.num)%f.order)
}

func (f *FieldElement) Subtract(other *FieldElement) *FieldElement {
	/*
		a, b element of the finite set, c = a - b, given b how can we find c,
		(b + c) % order = a, a - b => (a + (-b)) % order
	*/

	return f.Add(other.Negate())
}

func (f *FieldElement) Multiply(other *FieldElement) *FieldElement {
	f.checkOrder(other)

	// Arithmetic multiply over modular of the order
	return NewFieldElement(f.order, (f.num*other.num)%f.order)
}

func (f *FieldElement) Power(power int64) *FieldElement {
	// Arithmetic power over modular of the order
	return NewFieldElement(f.order, uint64(math.Pow(float64(f.num), float64(power)))%f.order)
}

func (f *FieldElement) ScalarMul(val uint64) *FieldElement {
	return NewFieldElement(f.order, (f.num*val)%f.order)
}
