package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bshefflette/fraction_calc/fraction"
)

var VALID_OPERATORS = []string{"+", "-", "*", "/"}

// var OPERAND_REGEXP = "^(\\-?\\d+)?_?(\\-?\\d+)?\\/?(\\-?\\d+?)?$"
// var OPERAND_REGEXP = `^(\-?\d+)?_?(\-?\d+)?\/?(\-?\d+?)?$`

// var OPERAND_REGEXP = "(\\-?\\d+)\\/(\\-?\\d+)$|^(\\-?\\d+)"

var OPERAND_REGEXP = `(?P<numerator>\-?\d+)\/(?P<denominator>\-?\d+)|^(?P<whole>\-?\d+)`

type FractionOperation struct {
	Operand1  Operand
	Operand2  Operand
	Operation string
	Result    Operand
}
type Operand struct {
	Numerator   string
	Denominator string
	Whole       string
}

func main() {

	var s string

	// only accept 1 argument for input string
	if len(os.Args) == 2 {
		if strings.ToLower(os.Args[1]) == "help" {
			PrintHelp()
			os.Exit(0)
		}
		s = os.Args[1]
	} else {
		fmt.Println("run program with argument \"help\" for instructions")
		// error if too many args
		if len(os.Args) > 2 {
			log.Fatalf("ERROR too many arguments (%d) [%s] in command. Only 1 allowed.", len(os.Args), os.Args)
		}
		os.Exit(1)
	}

	fo, err := ParseFractionOperationsString(s)
	if err != nil {
		log.Fatalf("ERROR parsing FractionOperation string [%s] %s", s, err)
	}
	err = fo.Process()
	if err != nil {
		log.Fatalf("ERROR processing FractionOperation: %s", err)
	}

	// output result
	fmt.Println(fo.Result.String())
}
func (fo *FractionOperation) Process() error {
	var err error
	var fr fraction.Fraction
	// process operation

	// o1 -> f1
	// o2 -> f2
	f1, err := fo.Operand1.ToFraction()
	if err != nil {
		return err
	}
	f2, err := fo.Operand2.ToFraction()
	if err != nil {
		return err
	}
	// fmt.Printf("f1: %+v\nf2:%+v\n", f1, f2)
	switch fo.Operation {
	case "+":
		fr = fraction.Add(f1, f2)
	case "-":
		fr = fraction.Subtract(f1, f2)
	case "*":
		fr = fraction.Multiply(f1, f2)
	case "/":
		fr = fraction.Divide(f1, f2)
	default:
		return fmt.Errorf("ERROR invalid operation sign [%s]", fo.Operation)
	}
	// fmt.Printf("fr: %+v\n", fr)
	fr.ReduceFraction()

	// use operation to determine function
	// fractFunc(f1,f2) -> res -> o3 (res)
	//
	fo.Result = FractionToOperand(fr)

	return err
}

func FractionToOperand(f fraction.Fraction) Operand {
	var res Operand
	var n, d int

	// normalize signs and reduce factors
	f.ReduceFraction()

	if f.Denominator == 0 {
		res.Numerator = strconv.Itoa(f.Numerator)
		res.Denominator = "0"
		return res
	}

	// if numerator is 0, return as whole number
	if f.Numerator == 0 {
		res.Whole = "0"
		return res
	}

	n, d = f.Numerator, f.Denominator
	// abs value n
	if n < 0 {
		n = -1 * n
	}

	// if abs(numerator) > denominator,
	// need to reduce for whole number
	if n >= d {
		// using f.Numerator/d will give w proper sign
		res.Whole = strconv.Itoa(f.Numerator / d)
		rem := n % d
		if rem > 0 {
			res.Numerator = strconv.Itoa(rem)
			res.Denominator = strconv.Itoa(d)
		}
	} else {
		res.Numerator = strconv.Itoa(f.Numerator)
		res.Denominator = strconv.Itoa(f.Denominator)
	}
	return res
}

/*
	String fractions must be in one of the following formats:
	leadingnegative_wholenumber_numerator/denominator
	"-3_1/4" "-1/8" or "-2"
	negative sign is optional, must be leading all other numbers
*/
func (o *Operand) ToFraction() (fraction.Fraction, error) {
	var res fraction.Fraction
	var err error
	err = o.Validate()
	if err != nil {
		return res, err
	}
	hasW, hasF := false, false

	var w, n, d int

	// calc whole value
	if o.Whole != "" {
		// hasW is used later to convert into fraction
		hasW = true
		w, err = strconv.Atoi(o.Whole)
		if err != nil {
			return res, err
		}
	}

	if o.Numerator != "" {
		// fraction must have a valid numerator
		// or not considered a fraction
		hasF = true
		n, err = strconv.Atoi(o.Numerator)
		if err != nil {
			return res, err
		}
		d, err = strconv.Atoi(o.Denominator)
		if err != nil {
			return res, err
		}
	}

	// if doesn't have fraction, use 0,1 for 0 value
	if !hasF {
		n, d = 0, 1
	}
	// convert n & d to fraction
	// initfraction checks for 0 denominator and
	// will reduce sign & factors if needed
	res, err = fraction.InitFraction(n, d)
	if err != nil {
		return res, err
	}

	// if whole number, add it to numerator by cross mult to denominator
	if hasW {
		if w < 0 && res.Numerator < 0 {
			// both negative, make positive
			w = -1 * w
			res.Numerator = -1 * res.Numerator
		} else if w < 0 {
			// only w is negative, make numerator negative too
			res.Numerator = -1 * res.Numerator
		} else if res.Numerator < 0 {
			// only numerator is negative, make w negative too
			w = -1 * w
		} else {
			// neither is negative, run as normal
		}
		res.Numerator = res.Numerator + (w * res.Denominator)
	}

	return res, err
}

func ParseFractionOperationsString(s string) (FractionOperation, error) {
	var res FractionOperation
	var err error
	fs := strings.Fields(s)
	if len(fs) != 3 {
		return res, fmt.Errorf("ERROR invalid input [%s], must be formatted like [operand operator operand] ", s)
	}
	o1, err := ParseOperandString(fs[0])
	if err != nil {
		return res, err
	}

	o2, err := ParseOperandString(fs[2])
	if err != nil {
		return res, err
	}

	operator := fs[1]
	err = ValidateOperator(operator)
	if err != nil {
		return res, err
	}

	res.Operand1 = o1
	res.Operand2 = o2
	res.Operation = operator
	return res, nil
}

func ParseOperandString(s string) (Operand, error) {
	var err error
	var res Operand

	re := regexp.MustCompile(OPERAND_REGEXP)
	if !re.Match([]byte(s)) {
		return res, fmt.Errorf("ERROR operand [%s] is bad format, should be like ['1', '1_3/8', or '9/6']", s)
	}

	groups := re.FindAllStringSubmatch(s, -1)

	// create map for consolidating regexp group values
	m := make(map[string]string)
	// loop through capture groups and capture names to store values
	for i := range groups {
		for j, v := range re.SubexpNames() {
			if groups[i][j] != "" {
				m[v] = groups[i][j]
			}
		}
	}

	// set operand
	res = Operand{Numerator: m["numerator"], Denominator: m["denominator"], Whole: m["whole"]}
	err = res.Validate()

	// fmt.Printf("res: %+v\n", res)
	return res, err

}
func (o *Operand) Validate() error {
	// operand must have non-0 denominator
	// operand must have num + den or whole

	if o.Numerator != "" && o.Denominator == "" {
		return fmt.Errorf("ERROR operand [%s] has valid numerator, missing valid denominator", o)
	}
	if o.Numerator == "" && o.Denominator != "" {
		return fmt.Errorf("ERROR operand [%s] has valid denominator, missing valid numerator", o)
	}
	if o.Numerator == "" && o.Denominator == "" && o.Whole == "" {
		return fmt.Errorf("ERROR operand [%s] has is missing all valid values", o)
	}
	return nil
}
func (o *Operand) String() string {
	// need to modify this to output differently based on operand value
	if o.Whole == "" && o.Numerator == "" {
		// no whole or fraction, print 0 as default state
		return "0"
	}
	if o.Whole == "" {
		// no whole, print fraction only
		return fmt.Sprintf("%s/%s", o.Numerator, o.Denominator)
	}
	if o.Numerator == "" {
		// no fraction, print only whole
		return o.Whole
	}
	return fmt.Sprintf("%s_%s/%s", o.Whole, o.Numerator, o.Denominator)
}
func ValidateOperator(s string) error {
	for _, v := range VALID_OPERATORS {
		if v == s {
			return nil
		}
	}
	return fmt.Errorf("ERROR no valid operator provided [%s]", s)
}
func PrintHelp() {
	fmt.Printf(`Must provide argument for fraction calculation:
./fraction_calc "2_3/8 + 1/9"
	-> 2_35/72
./fraction_calc "-2_3/-8 + -0_-1/-9"
	-> 2_35/72
./fraction_calc "-19/-8 - 1/-9"
	-> 2_35/72

- Format: "{operand}{spaces}{operation}{spaces}{operand}
	- Operand(s) can be whole number, fraction, or mixed number
		- Operand Formats:
			- Whole Number Format: 	"{integer}"
			- Fraction Format: 		"{integer}/{integer}"
			- Mixed Number Format: 	"{wholenumber}_{fraction}"
		- Fractions can be improper
		- All numbers within an operand must be integers (each can have 1 negative sign if needed)
		- Do not divide by 0 within an operand fraction, will error
	- Operation must be one of: ["+", "-", "*", "/"] 
		- (add, subtract, multiply, divide)
	- Space(s) must be whitespace
- Interesting Notes:
	- Whole number can be 0
	- Negative 0 will be treated as 0
	- Fractions that divide by 0 will error during evaluation
	- Fractions with 0 in numerator will always reduce to 0/1
		- 0/2 -> 0/1
		- 0/-2 -> 0/1
	- 0's will be ignored if in mixed number
	- You any input like "{operand} / 0" will reduce to 1/0
		- Unless it errors for another reason (e.g. 1/0 [because fraction in operand divides by 0])
`)
}
