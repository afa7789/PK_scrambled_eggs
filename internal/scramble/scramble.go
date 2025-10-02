package scramble

import (
	"fmt"
	"os"

	pkgscramble "github.com/afa7789/PK_scrambled_eggs/pkg/scramble"
)

func Scramble(
	intFlag int,
	stringFlag string,
	dateFlag string,
	debug bool,
) error {
	// Use pkg/scramble for scrambling
	scrambled, err := pkgscramble.Scramble(&intFlag, &stringFlag, &dateFlag)
	if err != nil {
		return fmt.Errorf("scramble failed: %w", err)
	}

	// Use pkg/scramble for unscrambling
	unscrambled, err := pkgscramble.Unscramble(&intFlag, &stringFlag, &dateFlag)
	if err != nil {
		return fmt.Errorf("unscramble failed: %w", err)
	}

	var output string
	if debug {
		output += fmt.Sprintf("Original input: \"%s\"\n", stringFlag)
		output += fmt.Sprintf("Scrambled: \"%s\"\n", scrambled)
		output += fmt.Sprintf("Unscrambled: \"%s\"\n\n", unscrambled)
	} else {
		output = fmt.Sprintf("Scrambled:\n%s\n\nUnscrambled:\n%s\n", scrambled, unscrambled)
	}

	// Write both scrambled and unscrambled to output.txt
	fileOutput := fmt.Sprintf("Scrambled:\n%s\n\nUnscrambled:\n%s\n", scrambled, unscrambled)

	// Write output to output.txt and print to screen
	fmt.Print(output)
	err = os.WriteFile("output.txt", []byte(fileOutput), 0644)
	if err != nil {
		fmt.Printf("Error writing to output.txt: %v\n", err)
		return err
	}

	return nil
}
