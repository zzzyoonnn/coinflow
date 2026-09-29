package elliptic_curve

import (
	"fmt"
	"math/big"
)

type OP_TYPE int

const (
	ADD OP_TYPE = iota
	SUB
	MUL
	DIV
	EXP
)

type Point struct {
	// coefficients of curve
	a *FieldElement
	b *FieldElement

	// x, y should be the point on the curve
	x *FieldElement
	y *FieldElement
}

func OpOnBig(x *FieldElement, y *FieldElement, scalar *big.Int, opType OP_TYPE) *FieldElement {
	switch opType {
	case ADD:
		return x.Add(y)
	case SUB:
		return x.Subtract(y)
	case MUL:
		if y != nil {
			return x.Multiply(y)
		}
		if scalar != nil {
			return x.ScalarMul(scalar)
		}
		panic("error in multiply")
	case DIV:
		return x.Divide(y)
	case EXP:
		if scalar == nil {
			panic("scalar should not be nil for EXP")
		}
		return x.Power(scalar)
	}

	panic("should not come to here")
}

func NewEllipticCurvePoint(x *FieldElement, y *FieldElement, a *FieldElement, b *FieldElement) *Point {
	if x == nil && y == nil {
		return &Point{
			x: x,
			y: y,
			a: a,
			b: b,
		}
	}
	left := OpOnBig(y, nil, big.NewInt(int64(2)), EXP)
	x3 := OpOnBig(x, nil, big.NewInt(int64(3)), EXP)
	ax := OpOnBig(a, x, nil, MUL)
	right := OpOnBig(OpOnBig(x3, ax, nil, ADD), b, nil, ADD)

	if left.EqualTo(right) != true {
		err := fmt.Sprintf("Point(%v, %v) is not on the curve with a:%v, b:%v\n", x, y, a, b)
		panic(err)
	}

	return &Point{
		x: x,
		y: y,
		a: a,
		b: b,
	}
}

func (p *Point) Add(other *Point) *Point {
	// check two points are on the same curve
	if p.a.EqualTo(other.a) != true || p.b.EqualTo(other.b) != true {
		panic("given two points are not on the same curve")
	}

	if p.x == nil {
		return other
	}

	if other.x == nil {
		return p
	}

	zero := NewFieldElement(p.x.order, big.NewInt(int64(0)))
	// points are on the vertical A(x, y) B(x, -y)
	if p.x.EqualTo(other.x) == true && OpOnBig(p.y, other.y, nil, ADD).EqualTo(zero) == true {
		return &Point{
			x: nil,
			y: nil,
			a: p.a,
			b: p.b,
		}
	}

	// find slope of line AB
	// x1 -> p.x, y1 -> p.y, x2 -> other.x, y2 -> other.y
	var numerator *FieldElement
	var denominator *FieldElement
	if p.x.EqualTo(other.x) == true && p.y.EqualTo(other.y) == true {
		// slope = (3*x^2 + a) / 2y
		xSquared := OpOnBig(p.x, nil, big.NewInt(int64(2)), EXP)
		threeXSquared := OpOnBig(xSquared, nil, big.NewInt(int64(3)), MUL)
		numerator = OpOnBig(threeXSquared, p.a, nil, ADD)

		// denominator: 2y
		denominator = OpOnBig(p.y, nil, big.NewInt(int64(2)), MUL)
	} else {
		numerator = OpOnBig(other.y, p.y, nil, SUB)   // (y2 - y1)
		denominator = OpOnBig(other.x, p.x, nil, SUB) // (x2 - x1)
	}

	// s = (y2 - y1) / (x2 - x1)
	slope := OpOnBig(numerator, denominator, nil, DIV)

	// s^2
	slopeSqrt := OpOnBig(slope, nil, big.NewInt(int64(2)), EXP)

	// x3 = s^2 - x1 - x2
	x3 := OpOnBig(OpOnBig(slopeSqrt, p.x, nil, SUB), other.x, nil, SUB)

	// x3 - x1
	x3Minusx1 := OpOnBig(x3, p.x, nil, SUB)

	// y3 = s(x3 - x1) + y1
	y3 := OpOnBig(OpOnBig(slope, x3Minusx1, nil, MUL), p.y, nil, ADD)

	// -y3
	minusY3 := OpOnBig(y3, nil, big.NewInt(int64(-1)), MUL)

	return &Point{
		x: x3,
		y: minusY3,
		a: p.a,
		b: p.b,
	}
}

func (p *Point) String() string {
	xString := "nil"
	yString := "nil"

	if p.x != nil {
		xString = p.x.String()
	}
	if p.y != nil {
		yString = p.y.String()
	}

	return fmt.Sprintf("(x:%s, y:%s, a:%s, b:%s)", xString, yString, p.a.String(), p.b.String())
}

func (p *Point) Equal(other *Point) bool {
	if p.a.EqualTo(other.a) == true && p.b.EqualTo(other.b) == true && p.x.EqualTo(other.x) == true && p.y.EqualTo(other.y) == true {
		return true
	}

	return false
}

func (p *Point) NotEqual(other *Point) bool {
	if p.a.EqualTo(other.a) != true || p.b.EqualTo(other.b) != true || p.x.EqualTo(other.x) != true || p.y.EqualTo(other.y) != true {
		return true
	}

	return false
}
