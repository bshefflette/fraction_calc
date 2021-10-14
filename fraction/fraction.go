package fraction

import (
	"fmt"
)

type Fraction struct {
	Numerator   int
	Denominator int
}

func Add(f1, f2 Fraction) Fraction {
	res := Fraction{}

	n1, n2, d1, d2 := f1.Numerator, f2.Numerator, f1.Denominator, f2.Denominator
	if d1 == d2 {
		// simply add n1 + n2
		res.Numerator = n1 + n2
		res.Denominator = d2
	} else {
		//
		res.Numerator = (n1 * d2) + (n2 * d1)
		res.Denominator = d1 * d2
	}
	return res
}
func Subtract(f1, f2 Fraction) Fraction {
	// flip second operand sign then add

	f2.Numerator = -1 * f2.Numerator

	return Add(f1, f2)
}
func Multiply(f1, f2 Fraction) Fraction {
	// num1 * num2 && den1 * den2
	res := Fraction{}
	res.Numerator = f1.Numerator * f2.Numerator
	res.Denominator = f1.Denominator * f2.Denominator
	return res
}
func Divide(f1, f2 Fraction) Fraction {
	// flip  f2 num & den, mult
	f2.Numerator, f2.Denominator = f2.Denominator, f2.Numerator
	f2.ReduceSign()
	return Multiply(f1, f2)
}

func InitFraction(num int, den int) (Fraction, error) {
	res := Fraction{num, den}
	if den == 0 {
		return res, fmt.Errorf("ERROR denominator cannot be 0")
	}
	res.ReduceFraction()
	return res, nil
}

func (f *Fraction) ReduceFraction() {
	f.ReduceSign()
	f.ReduceFactors()
}

func (f *Fraction) ReduceSign() {
	if f.Denominator < 0 {
		// denominator is negative, not numerator
		// or both are negative
		// flip both signs
		f.Denominator = -1 * f.Denominator
		f.Numerator = -1 * f.Numerator
	}
}

func Gcf(a, b int) int {
	// non-recursive alternative:
	if a < 0 {
		a = -1 * a
	}
	if b < 0 {
		b = -1 * b
	}
	for b != 0 {
		c := b
		b = a % b
		a = c
	}
	return a
}
func (f *Fraction) ReduceFactors() {

	gcf := Gcf(f.Numerator, f.Denominator)
	if gcf == 0 {
		return
	}
	f.Numerator = f.Numerator / gcf
	f.Denominator = f.Denominator / gcf
}
