package natural_const

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"
)

// Constant
type Constant struct {
	Digits           []int
	number_of_digits int
	decimalPos       int // Track the decimal point position
}

// NaturalConstant
type NaturalConstant struct {
	base int
	Constant
}

// CustomConstant
type CustomConstant struct {
	Constant
	BaseNaturalConstant *NaturalConstant
	date                time.Time
	floatDate           big.Float
}

// String returns the digits of the custom constant as a string.
func (c *Constant) String() string {
	var sb strings.Builder
	for i, digit := range c.Digits {
		if i == c.decimalPos { // Add decimal point
			sb.WriteString(".")
		}
		sb.WriteString(fmt.Sprintf("%d", digit))
	}
	return sb.String()
}

func NewConstant(digits []int) *Constant {
	return &Constant{Digits: digits, number_of_digits: len(digits)}
}

func NewNaturalConstant(baseInt int) *NaturalConstant {
	return &NaturalConstant{
		base:     baseInt,
		Constant: Constant{Digits: []int{}, number_of_digits: 0, decimalPos: 1},
	}
}

// NewCustomConstant creates a new CustomConstant.
func NewCustomConstant(baseInt int, date time.Time) *CustomConstant {
	parseFloat := fmt.Sprintf(".%02d%02d%04d", date.Month(), date.Day(), date.Year())
	floatDate, _, err := big.ParseFloat(parseFloat, 10, 10000, big.ToNearestEven)
	if err != nil {
		floatDate = big.NewFloat(0)
	}

	nc := NewNaturalConstant(baseInt)
	return &CustomConstant{
		BaseNaturalConstant: nc,
		date:                date,
		floatDate:           *floatDate,
	}
}

// CalculateCustomConstant calculates the custom constant by dividing the natural constant by the date.
// amount is the number of digits to calculate.
func (cc *CustomConstant) CalculateCustomConstant(amount int) {
	cc.BaseNaturalConstant.CalculateDigits(amount)

	ncValue := cc.BaseNaturalConstant.toBigFloat(amount)
	// dateValue := big.NewFloat(cc.floatDate).SetPrec(uint(amount + 2))
	customValue := new(big.Float).Quo(ncValue, &cc.floatDate)
	cc.populateDigits(customValue, amount)
}

// toBigFloat converts the digits of the natural constant to a big.Float.
func (nc *NaturalConstant) toBigFloat(amount int) *big.Float {
	str := nc.String()
	value, _, _ := big.ParseFloat(str, 10, uint(amount+2), big.ToNearestEven)
	return value
}

// Name returns the name of the constant based on an integer.
func (nc NaturalConstant) Name() string {
	switch nc.base {
	case 0:
		return "PI - π"
	case 1:
		return "Euler's number - e"
	case 2:
		return "Golden ratio (phi) - ϕ"
	case 3:
		return "Natural logarithm of 2 - ln(2)"
	case 4:
		return "Euler's constant - γ"
	case 5:
		return "Pythagoras' constant - √2"
	default:
		return "Unknown Constant"
	}
}

// CalculateDigits calculates the digits of the natural constant up to the specified amount dynamically.
func (nc *NaturalConstant) CalculateDigits(amount int) {
	switch nc.base {
	case 0:
		nc.calculatePi(amount)
	case 1:
		nc.calculateENumber(amount)
	case 2:
		nc.calculatePhi(amount)
	case 3:
		nc.calculateLn2(amount)
	case 4:
		nc.calculateGamma(amount)
	case 5:
		nc.calculateSqrt2(amount)
	}
}

func (nc *NaturalConstant) calculatePi(amount int) {
	// Add extra digits to account for rounding
	precision := uint(amount + 10)
	pi := big.NewFloat(0).SetPrec(precision)
	four := big.NewFloat(4).SetPrec(precision)
	one := big.NewFloat(1).SetPrec(precision)
	two := big.NewFloat(2).SetPrec(precision)
	sixteen := big.NewFloat(16).SetPrec(precision)

	// Initialize the denominator to 1 (for 16^0)
	denominator := new(big.Float).SetPrec(precision).SetFloat64(1)

	// Implement BBP formula
	for k := int64(0); k <= int64(amount); k++ {
		// Initialize term for each iteration
		term := big.NewFloat(0).SetPrec(precision)

		// term = 4 / (8k + 1) - 2 / (8k + 4) - 1 / (8k + 5) - 1 / (8k + 6)
		term.Add(term, new(big.Float).Quo(four, big.NewFloat(float64(8*k+1))))
		term.Sub(term, new(big.Float).Quo(two, big.NewFloat(float64(8*k+4))))
		term.Sub(term, new(big.Float).Quo(one, big.NewFloat(float64(8*k+5))))
		term.Sub(term, new(big.Float).Quo(one, big.NewFloat(float64(8*k+6))))

		// Divide term by 16^k (incrementally multiply denominator by 16)
		if k > 0 {
			denominator.Mul(denominator, sixteen) // Multiply by 16 each time
		}

		// Divide term by the updated denominator (16^k)
		term.Quo(term, denominator)

		// Add term to pi
		pi.Add(pi, term)
	}

	// Convert pi to a string and populate digits
	nc.populateDigits(pi, amount)
}

// calculateE calculates Euler's number e using a series expansion.
func (nc *NaturalConstant) calculateENumber(amount int) {
	precision := uint(amount + 2)           // Add extra digits for rounding
	e := big.NewFloat(2).SetPrec(precision) // e starts at 2 (1 + 1/1!)
	factorial := big.NewFloat(1).SetPrec(precision)

	for i := 2; i < amount+2; i++ {
		factorial.Mul(factorial, big.NewFloat(float64(i)))
		term := new(big.Float).Quo(big.NewFloat(1), factorial)
		e.Add(e, term)
	}

	// Convert e to a string and populate digits
	nc.populateDigits(e, amount)
}

// calculatePhi calculates the golden ratio (ϕ) to the specified number of digits.
func (nc *NaturalConstant) calculatePhi(amount int) {
	precision := uint(amount + 2) // Add extra digits to account for rounding
	one := big.NewFloat(1).SetPrec(precision)
	two := big.NewFloat(2).SetPrec(precision)
	five := big.NewFloat(5).SetPrec(precision)

	// Calculate sqrt(5)
	sqrt5 := big.NewFloat(0).SetPrec(precision).Sqrt(five)

	// Calculate phi = (1 + sqrt(5)) / 2
	phi := big.NewFloat(0).SetPrec(precision)
	phi.Add(one, sqrt5)
	phi.Quo(phi, two)

	// Convert phi to a string and populate digits
	nc.populateDigits(phi, amount)
}

// calculateLn2 calculates the natural logarithm of 2 to the specified precision.
func (nc *NaturalConstant) calculateLn2(amount int) {
	// Set the precision to the desired number of digits plus a few extra for accuracy
	precision := uint(amount + 2)
	ln2 := new(big.Float).SetPrec(precision).SetFloat64(math.Ln2)

	// Convert the big.Float to a string and extract the digits
	ln2Str := ln2.Text('f', amount)
	nc.decimalPos = strings.Index(ln2Str, ".")   // Track the decimal point position
	ln2Str = strings.ReplaceAll(ln2Str, ".", "") // Remove the decimal point for digit extraction

	nc.Digits = make([]int, len(ln2Str))
	for i, char := range ln2Str {
		nc.Digits[i] = int(char - '0')
	}
	nc.number_of_digits = len(nc.Digits)
}

// calculateGamma calculates Euler's constant (gamma) to the specified precision.
func (nc *NaturalConstant) calculateGamma(amount int) {
	// Set the precision to the desired number of digits plus a few extra for accuracy
	precision := uint(amount + 2)
	gamma := new(big.Float).SetPrec(precision).SetFloat64(0.57721566490153286060) // Approximation of Euler's constant

	// Convert the big.Float to a string and extract the digits
	gammaStr := gamma.Text('f', amount)
	nc.decimalPos = strings.Index(gammaStr, ".")     // Track the decimal point position
	gammaStr = strings.ReplaceAll(gammaStr, ".", "") // Remove the decimal point for digit extraction

	nc.Digits = make([]int, len(gammaStr))
	for i, char := range gammaStr {
		nc.Digits[i] = int(char - '0')
	}
	nc.number_of_digits = len(nc.Digits)
}

// calculateSqrt2 calculates the square root of 2 to the specified precision.
func (nc *NaturalConstant) calculateSqrt2(amount int) {
	// Set the precision to the desired number of digits plus a few extra for accuracy
	precision := uint(amount + 2)
	sqrt2 := new(big.Float).SetPrec(precision).SetFloat64(math.Sqrt2)

	// Convert the big.Float to a string and extract the digits
	sqrt2Str := sqrt2.Text('f', amount)
	nc.decimalPos = strings.Index(sqrt2Str, ".")     // Track the decimal point position
	sqrt2Str = strings.ReplaceAll(sqrt2Str, ".", "") // Remove the decimal point for digit extraction

	nc.Digits = make([]int, len(sqrt2Str))
	for i, char := range sqrt2Str {
		nc.Digits[i] = int(char - '0')
	}
	nc.number_of_digits = len(nc.Digits)
}

// populateDigits converts a *big.Float to digits and tracks the decimal point.
func (c *Constant) populateDigits(value *big.Float, amount int) {
	str := value.Text('f', amount)         // Get string representation with 'amount' digits after the decimal
	c.decimalPos = strings.Index(str, ".") // Track the decimal point position
	str = strings.ReplaceAll(str, ".", "") // Remove the decimal point for digit extraction

	c.Digits = make([]int, len(str))
	for i, ch := range str {
		c.Digits[i] = int(ch - '0')
	}
	c.number_of_digits = len(c.Digits)
}
