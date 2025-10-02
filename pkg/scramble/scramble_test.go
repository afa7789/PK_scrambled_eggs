package scramble

import (
	"testing"
)

func TestMultilineScrambleUnscramble(t *testing.T) {
	input := "alpha\nbeta\ngamma\ndelta\nepsilon\nzeta\neta\ntheta\niota\nkappa\nlambda\nmu\nnu\nxi\nomicron\npi\nrho\nsigma\ntau\nupsilon\nphi\nchi\npsi\nomega\n"
	intFlag := 3
	dateFlag := "09-11-2001"

	output, err := Scramble(&intFlag, &input, &dateFlag)
	if err != nil {
		t.Fatalf("Scramble failed: %v", err)
	}

	unscrambled, err := Unscramble(&intFlag, &output, &dateFlag)
	if err != nil {
		t.Fatalf("Unscramble failed: %v", err)
	}

	if input != unscrambled {
		t.Errorf("Unscrambled string does not match original.\nOriginal: %q\nUnscrambled: %q", input, unscrambled)
	}

	if input == unscrambled {
		t.Logf("✓ Verification successful\nunscrambled:\n%s\n\ninput:\n%s", unscrambled, input)
	} else {
		t.Errorf("✗ Verification failed")
	}
}
