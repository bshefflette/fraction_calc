package main

import "testing"

func TestToFraction(t *testing.T) {
	o1 := Operand{Numerator: "3", Denominator: "7", Whole: "5"}
	f1, err := o1.ToFraction()
	if err != nil {
		t.Fatalf("ERROR operand o1: [%s], f1: [%d/%d], err: %s", o1, f1.Numerator, f1.Denominator, err)
	}
	if f1.Numerator != 38 && f1.Denominator != 7 {
		t.Fatalf("ERROR f1: [%d/%d] is not 38/7 ", f1.Numerator, f1.Denominator)
	}

}
func TestFull(t *testing.T) {
	// whole * whole
	fo1, err := ParseFractionOperationsString("-1 * -2")
	if err != nil {
		t.Fatalf("ERROR fo1: [%+v] %s", fo1, err)
	}
	err = fo1.Process()
	if err != nil {
		t.Fatalf("ERROR fo1: [%+v] %s", fo1, err)
	}
	want := "2"
	if fo1.Result.String() != "2" {
		t.Fatalf("ERROR result [%s] does not match expected value [%s]", fo1.Result.String(), want)
	}

	// mixed + mixed
	fo2, err := ParseFractionOperationsString("-100_1/2 + -0_2/3")
	if err != nil {
		t.Fatalf("ERROR fo2: [%+v] %s", fo2, err)
	}
	err = fo2.Process()
	if err != nil {
		t.Fatalf("ERROR fo2: [%+v] %s", fo2, err)
	}
	want2 := "-99_5/6"
	if fo2.Result.String() != want2 {
		t.Fatalf("ERROR result [%s] does not match expected value [%s]", fo2.Result.String(), want2)
	}

	// 0 / 0
	fo3, err := ParseFractionOperationsString("0 / -0")
	if err != nil {
		t.Fatalf("ERROR fo3: [%+v] %s", fo3, err)
	}
	err = fo3.Process()
	if err != nil {
		t.Fatalf("ERROR fo3: [%+v] %s", fo3, err)
	}
	want3 := "0/0"
	if fo3.Result.String() != want3 {
		t.Fatalf("ERROR result [%s] does not match expected value [%s]", fo3.Result.String(), want3)
	}

	// 0 / 1
	fo4, err := ParseFractionOperationsString("0 / -1")
	if err != nil {
		t.Fatalf("ERROR fo4: [%+v] %s", fo4, err)
	}
	err = fo4.Process()
	if err != nil {
		t.Fatalf("ERROR fo4: [%+v] %s", fo4, err)
	}
	want4 := "0"
	if fo4.Result.String() != want4 {
		t.Fatalf("ERROR result [%s] does not match expected value [%s]", fo4.Result.String(), want4)
	}

	// fract - whole
	fo5, err := ParseFractionOperationsString("2/3 - -1")
	if err != nil {
		t.Fatalf("ERROR fo5: [%+v] %s", fo5, err)
	}
	err = fo5.Process()
	if err != nil {
		t.Fatalf("ERROR fo5: [%+v] %s", fo5, err)
	}
	want5 := "1_2/3"
	if fo5.Result.String() != want5 {
		t.Fatalf("ERROR result [%s] does not match expected value [%s]", fo5.Result.String(), want5)
	}

	// improper fract - fract
	//  ((4257 + 77)  / 301)  -> (4334/301)
	fo6, err := ParseFractionOperationsString("99/7 - -11/43")
	if err != nil {
		t.Fatalf("ERROR fo6: [%+v] %s", fo6, err)
	}
	err = fo6.Process()
	if err != nil {
		t.Fatalf("ERROR fo6: [%+v] %s", fo6, err)
	}
	want6 := "14_120/301"
	if fo6.Result.String() != want6 {
		t.Fatalf("ERROR result [%s] does not match expected value [%s]", fo6.Result.String(), want6)
	}

	// improper whole fract - fract with weird negative signs
	// -11/3 - 1/2 -> -22/6 - 3/6 -> -25/6 -> -4_1/6
	fo7, err := ParseFractionOperationsString("-2_-5/-3 - -1/-2")
	if err != nil {
		t.Fatalf("ERROR fo7: [%+v] %s", fo7, err)
	}
	err = fo7.Process()
	if err != nil {
		t.Fatalf("ERROR fo7: [%+v] %s", fo7, err)
	}
	want7 := "-4_1/6"
	if fo7.Result.String() != want7 {
		t.Fatalf("ERROR result [%s] does not match expected value [%s]", fo7.Result.String(), want7)
	}

}
