package natural_const

import (
	"fmt"
	"math/big"
	"testing"
	"time"
)

func TestNewCustomConstant(t *testing.T) {
	// t.Run("", func(t *testing.T) {
	date := time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC)
	cc := NewCustomConstant(0, date) // Using PI as the base natural constant

	if cc.floatDate.String() != "0.10052023" {
		t.Errorf("Expected floatDate to be 0.05102023, got %s", cc.floatDate.String())
	}
	// })
}

func TestCalculateCustomConstant(t *testing.T) {

	// custom amount, not yet good enough due to the fact that the float is not precise enough

	date := time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC)
	cc := NewCustomConstant(0, date) // Using PI as the base natural constant

	cc.CalculateCustomConstant(100) // Calculate 10 digits of the custom constant
	// used wolfram to check: pi/ 0.10052023
	// wolfram with first 100 digits of pi/ divided by 0.10052023
	// 31.2533372992659610753242743602904274542870232089855523995029464228929579323
	expected := "31.2533372"
	// 31.2533372
	// 0.31579362...

	result := cc.String()[:10] // Get the first 8 characters (including "0.")
	fmt.Printf("result %s", result)

	if result != expected {
		t.Errorf("Expected custom constant to start with %s, got %s", expected, result)
	}
}

func TestNaturalConstantCalculateDigits(t *testing.T) {

	tests := []struct {
		name     string
		base     int
		expected string
	}{
		{"PI", 0, "3.1415926535"},                     // PI
		{"Euler's number", 1, "2.7182818284"},         // Euler's number (e)
		{"Golden Ratio", 2, "1.6180339887"},           // Golden ratio (phi)
		{"Natural logarithm of 2", 3, "0.6931471805"}, // Natural logarithm of 2 (ln(2))
		{"Euler's constant", 4, "0.5772156649"},       // Euler's constant (gamma)
		{"Square root of 2", 5, "1.4142135623"},       // Square root of 2 (sqrt(2))
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// t.func() // TestNaturalConstantCalculateDigits
			nc := NewNaturalConstant(test.base)
			nc.CalculateDigits(100)
			result := nc.String()[:12] // Get the first 12 characters

			if result != test.expected {
				t.Errorf("Expected constant to start with %s, got %s", test.expected, result)
			}
		})
	}

}

func TestNaturalConstantToBigFloat(t *testing.T) {
	nc := NewNaturalConstant(0) // Using PI as the base natural constant
	nc.CalculateDigits(100)     // Calculate 10 digits of PI

	// expected := "3.1415926535"
	// result := nc.String()[:12] // Get the first 12 characters (including "3.")
	result := nc.toBigFloat(100)

	expected, _, _ := big.ParseFloat("3.1415926535897932384626433832794966345032460589825657938750880601068773967377012468205066397786140442", 10, 100, big.ToNearestEven)

	if response := result.Cmp(expected); response != 0 {
		t.Errorf("Expected PI to start with %s, got %s", expected.Text('f', 10), result.Text('f', 10))
	}
}
