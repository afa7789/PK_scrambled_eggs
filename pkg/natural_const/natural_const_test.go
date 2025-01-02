package natural_const

import (
	"fmt"
	"math/big"
	"testing"
	"time"
)

func TestNewCustomConstant(t *testing.T) {
	date := time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC)
	cc := NewCustomConstant(0, date) // Using PI as the base natural constant

	if cc.floatDate.String() != "0.10052023" {
		t.Errorf("Expected floatDate to be 0.05102023, got %s", cc.floatDate.String())
	}
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
	nc := NewNaturalConstant(0) // Using PI as the base natural constant
	nc.CalculateDigits(100)     // Calculate 10 digits of PI

	expected := "3.1415926535"
	result := nc.String()[:12] // Get the first 12 characters (including "3.")

	if result != expected {
		t.Errorf("Expected PI to start with %s, got %s", expected, result)
	}

	// check other natural constants
	// from woflram alpha: 2.71828182845904523536028747135266249775724709369995957496696762772407663035...
	nc = NewNaturalConstant(1)
	nc.CalculateDigits(100)
	fmt.Println(nc.String())
	// 2.7182818284590452353602874713536013977993050833954989575140352708881863463830086402595043182373046875

	// from wolfram 1.61803398874989484820458683436563811772030917980576286213544862270526046281...
	nc = NewNaturalConstant(2)
	nc.CalculateDigits(100)
	fmt.Println(nc.String())
	// 1.6180339887498948482045868343657895377906517771085788355350451594993899107066681608557701110839843750

	// from wolfram 0.693147180559945309417232121458176568075500134360255254120680009493393622
	nc = NewNaturalConstant(3)
	nc.CalculateDigits(100)
	fmt.Println(nc.String())
	// 6.931471805599452862267639829951804131269454956054687500000000000000000000000000000000000000000000000

	// from wolfram 0.57721566490153286060651209008240243104215933593992359880576723488486772677...
	nc = NewNaturalConstant(4)
	nc.CalculateDigits(100)
	fmt.Println(nc.String())
	// 5.772156649015328655494272425130475312471389770507812500000000000000000000000000000000000000000000000

	// wolfram alpha: 1.41421356237309504880168872420969807856967187537694807317667973799073248
	nc = NewNaturalConstant(5)
	nc.CalculateDigits(100)
	fmt.Println(nc.String())
	// 4.142135623730951454746218587388284504413604736328125000000000000000000000000000000000000000000000000

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
