package fraction

import (
	"testing"
)

func TestAdd(t *testing.T) {
	f1, _ := InitFraction(1, 2)
	f2, _ := InitFraction(4, 9)
	f3, _ := InitFraction(-1, 7)
	f4, _ := InitFraction(3, -7)
	f5, _ := InitFraction(23, 5)
	f6, _ := InitFraction(0, 4)
	t1 := Add(f1, f2)
	if t1.Numerator != 17 || t1.Denominator != 18 {
		t.Fatalf("t1 Add failed: %+v\n", t1)
	}
	t2 := Add(f1, f3)
	if t2.Numerator != 5 || t2.Denominator != 14 {
		t.Fatalf("t2 Add failed: %+v\n", t2)
	}
	t3 := Add(f1, f4)
	if t3.Numerator != 1 || t3.Denominator != 14 {
		t.Fatalf("t3 Add failed: %+v\n", t3)
	}
	t4 := Add(f3, f4)
	if t4.Numerator != -4 || t4.Denominator != 7 {
		t.Fatalf("t4 Add failed: %+v\n", t4)
	}
	t5 := Add(f1, f5)
	if t5.Numerator != 51 || t5.Denominator != 10 {
		t.Fatalf("t5 Add failed: %+v\n", t5)
	}
	t6 := Add(f1, f6)
	if t6.Numerator != 1 || t6.Denominator != 2 {
		t.Fatalf("t6 Add failed: %+v\n", t6)
	}
}
func TestSubtract(t *testing.T) {
	f1, _ := InitFraction(1, 2)
	f2, _ := InitFraction(4, 9)
	f3, _ := InitFraction(-1, 7)
	f4, _ := InitFraction(3, -7)
	f5, _ := InitFraction(23, 5)
	f6, _ := InitFraction(0, 4)
	// fmt.Println(fs, fsl)
	t1 := Subtract(f1, f2)
	if t1.Numerator != 1 || t1.Denominator != 18 {
		t.Fatalf("t1 Subtract failed: %+v\n", t1)
	}
	t2 := Subtract(f1, f3)
	if t2.Numerator != 9 || t2.Denominator != 14 {
		t.Fatalf("t2 Subtract failed: %+v\n", t2)
	}
	t3 := Subtract(f1, f4)
	if t3.Numerator != 13 || t3.Denominator != 14 {
		t.Fatalf("t3 Subtract failed: %+v\n", t3)
	}
	t4 := Subtract(f3, f4)
	if t4.Numerator != 2 || t4.Denominator != 7 {
		t.Fatalf("t4 Subtract failed: %+v\n", t4)
	}
	t5 := Subtract(f1, f5)
	if t5.Numerator != -41 || t5.Denominator != 10 {
		t.Fatalf("t5 Subtract failed: %+v\n", t5)
	}
	t6 := Subtract(f1, f6)
	if t6.Numerator != 1 || t6.Denominator != 2 {
		t.Fatalf("t6 Subtract failed: %+v\n", t6)
	}
}
func TestMultiply(t *testing.T) {
	f1, _ := InitFraction(1, 2)
	f2, _ := InitFraction(4, 9)
	f3, _ := InitFraction(-1, 7)
	f4, _ := InitFraction(3, -7)
	f5, _ := InitFraction(23, 5)
	f6, _ := InitFraction(0, 4)
	// fmt.Println(fs, fsl)
	t1 := Multiply(f1, f2)
	if t1.Numerator != 4 || t1.Denominator != 18 {
		t.Fatalf("t1 Multiply failed: %+v\n", t1)
	}
	t2 := Multiply(f1, f3)
	if t2.Numerator != -1 || t2.Denominator != 14 {
		t.Fatalf("t2 Multiply failed: %+v\n", t2)
	}
	t3 := Multiply(f1, f4)
	if t3.Numerator != -3 || t3.Denominator != 14 {
		t.Fatalf("t3 Multiply failed: %+v\n", t3)
	}
	t4 := Multiply(f3, f4)
	if t4.Numerator != 3 || t4.Denominator != 49 {
		t.Fatalf("t4 Multiply failed: %+v\n", t4)
	}
	t5 := Multiply(f1, f5)
	if t5.Numerator != 23 || t5.Denominator != 10 {
		t.Fatalf("t5 Multiply failed: %+v\n", t5)
	}
	t6 := Multiply(f1, f6)
	if t6.Numerator != 0 || t6.Denominator != 2 {
		t.Fatalf("t6 Multiply failed: %+v\n", t6)
	}
}
func TestDivide(t *testing.T) {
	f1, _ := InitFraction(1, 2)
	f2, _ := InitFraction(4, 9)
	f3, _ := InitFraction(-1, 7)
	f4, _ := InitFraction(3, -7)
	f5, _ := InitFraction(23, 5)
	f6, _ := InitFraction(0, 4)
	// fmt.Println(fs, fsl)
	t1 := Divide(f1, f2)
	if t1.Numerator != 9 || t1.Denominator != 8 {
		t.Fatalf("t1 Divide failed: %+v\n", t1)
	}
	t2 := Divide(f1, f3)
	if t2.Numerator != -7 || t2.Denominator != 2 {
		t.Fatalf("t2 Divide failed: %+v\n", t2)
	}
	t3 := Divide(f1, f4)
	if t3.Numerator != -7 || t3.Denominator != 6 {
		t.Fatalf("t3 Divide failed: %+v\n", t3)
	}
	t4 := Divide(f3, f4)
	if t4.Numerator != 7 || t4.Denominator != 21 {
		t.Fatalf("t4 Divide failed: %+v\n", t4)
	}
	t5 := Divide(f1, f5)
	if t5.Numerator != 5 || t5.Denominator != 46 {
		t.Fatalf("t5 Divide failed: %+v\n", t5)
	}
	t6 := Divide(f1, f6)
	if t6.Numerator != 1 || t6.Denominator != 0 {
		t.Fatalf("t6 Divide failed: %+v\n", t6)
	}
}
func TestGcf(t *testing.T) {
	num1, num2 := 4, 9
	gcf := Gcf(num1, num2)
	// fmt.Printf("gcf is: %d\n", gcf)
	if gcf != 1 {
		t.Fatalf(`gcf failed because: %d`, gcf)
	}

	num3, num4 := 6, 9
	gcf = Gcf(num3, num4)
	if gcf != 3 {
		t.Fatalf(`gcf failed because: %d`, gcf)
	}
	num5, num6 := 144, 256
	gcf = Gcf(num5, num6)
	if gcf != 16 {
		t.Fatalf(`gcf failed because: %d`, gcf)
	}
}

func TestReduceFactors(t *testing.T) {
	f1 := &Fraction{22, 42}
	f1.ReduceFactors()
	if f1.Numerator != 11 || f1.Denominator != 21 {
		t.Fatalf("reduce failed: %d / %d", f1.Numerator, f1.Denominator)
	}

	f2 := &Fraction{22, 6}
	f2.ReduceFactors()
	if f2.Numerator != 11 || f2.Denominator != 3 {
		t.Fatalf("reduce failed: %d / %d", f2.Numerator, f2.Denominator)
	}

	f3 := &Fraction{22, 1}
	f3.ReduceFactors()
	if f3.Numerator != 22 || f3.Denominator != 1 {
		t.Fatalf("reduce failed: %d / %d", f3.Numerator, f3.Denominator)
	}

	f4 := &Fraction{22, 3}
	f4.ReduceFactors()
	if f4.Numerator != 22 || f4.Denominator != 3 {
		t.Fatalf("reduce failed: %d / %d", f4.Numerator, f4.Denominator)
	}
	f5 := &Fraction{22, -3}
	f5.ReduceFactors()
	if f5.Numerator != 22 || f5.Denominator != -3 {
		t.Fatalf("reduce failed: %d / %d", f5.Numerator, f5.Denominator)
	}
	f6 := &Fraction{-22, 3}
	f6.ReduceFactors()
	if f6.Numerator != -22 || f6.Denominator != 3 {
		t.Fatalf("reduce failed: %d / %d", f6.Numerator, f6.Denominator)
	}
	f7 := &Fraction{0, 3}
	f7.ReduceFactors()
	if f7.Numerator != 0 || f7.Denominator != 1 {
		t.Fatalf("reduce failed: %d / %d", f7.Numerator, f7.Denominator)
	}
	f8 := &Fraction{3, 0}
	f8.ReduceFactors()
	if f8.Numerator != 1 || f8.Denominator != 0 {
		t.Fatalf("reduce failed: %d / %d", f8.Numerator, f8.Denominator)
	}
	f9 := &Fraction{-3, 0}
	f9.ReduceFactors()
	if f9.Numerator != -1 || f9.Denominator != 0 {
		t.Fatalf("reduce failed: %d / %d", f9.Numerator, f9.Denominator)
	}
	f10 := &Fraction{-3, -3}
	f10.ReduceFactors()
	if f10.Numerator != -1 || f10.Denominator != -1 {
		t.Fatalf("reduce failed: %d / %d", f10.Numerator, f10.Denominator)
	}

}
func TestReduceFraction(t *testing.T) {
	f1 := &Fraction{22, 42}
	f1.ReduceFraction()
	if f1.Numerator != 11 || f1.Denominator != 21 {
		t.Fatalf("reduce failed: %d / %d", f1.Numerator, f1.Denominator)
	}

	f2 := &Fraction{22, 6}
	f2.ReduceFraction()
	if f2.Numerator != 11 || f2.Denominator != 3 {
		t.Fatalf("reduce failed: %d / %d", f2.Numerator, f2.Denominator)
	}

	f3 := &Fraction{22, 1}
	f3.ReduceFraction()
	if f3.Numerator != 22 || f3.Denominator != 1 {
		t.Fatalf("reduce failed: %d / %d", f3.Numerator, f3.Denominator)
	}

	f4 := &Fraction{22, 3}
	f4.ReduceFraction()
	if f4.Numerator != 22 || f4.Denominator != 3 {
		t.Fatalf("reduce failed: %d / %d", f4.Numerator, f4.Denominator)
	}
	f5 := &Fraction{22, -3}
	f5.ReduceFraction()
	if f5.Numerator != -22 || f5.Denominator != 3 {
		t.Fatalf("reduce failed: %d / %d", f5.Numerator, f5.Denominator)
	}
	f6 := &Fraction{-22, 3}
	f6.ReduceFraction()
	if f6.Numerator != -22 || f6.Denominator != 3 {
		t.Fatalf("reduce failed: %d / %d", f6.Numerator, f6.Denominator)
	}
	f7 := &Fraction{0, 3}
	f7.ReduceFraction()
	if f7.Numerator != 0 || f7.Denominator != 1 {
		t.Fatalf("reduce failed: %d / %d", f7.Numerator, f7.Denominator)
	}
	f8 := &Fraction{3, 0}
	f8.ReduceFraction()
	if f8.Numerator != 1 || f8.Denominator != 0 {
		t.Fatalf("reduce failed: %d / %d", f8.Numerator, f8.Denominator)
	}
	f9 := &Fraction{-3, 0}
	f9.ReduceFraction()
	if f9.Numerator != -1 || f9.Denominator != 0 {
		t.Fatalf("reduce failed: %d / %d", f9.Numerator, f9.Denominator)
	}
	f10 := &Fraction{-3, -3}
	f10.ReduceFraction()
	if f10.Numerator != 1 || f10.Denominator != 1 {
		t.Fatalf("reduce failed: %d / %d", f10.Numerator, f10.Denominator)
	}

}
func TestReduceSign(t *testing.T) {
	f1 := Fraction{1, -3}
	f1.ReduceSign()
	if f1.Numerator != -1 || f1.Denominator != 3 {
		t.Fatalf("reduce failed: %d / %d", f1.Numerator, f1.Denominator)
	}
	f2 := Fraction{-1, -3}
	f2.ReduceSign()
	if f2.Numerator != 1 || f2.Denominator != 3 {
		t.Fatalf("reduce failed: %d / %d", f2.Numerator, f2.Denominator)
	}
	f3 := Fraction{1, 3}
	f3.ReduceSign()
	if f3.Numerator != 1 || f3.Denominator != 3 {
		t.Fatalf("reduce failed: %d / %d", f3.Numerator, f3.Denominator)
	}
	f4 := Fraction{-1, 3}
	f4.ReduceSign()
	if f4.Numerator != -1 || f4.Denominator != 3 {
		t.Fatalf("reduce failed: %d / %d", f4.Numerator, f4.Denominator)
	}
}
