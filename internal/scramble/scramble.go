package scramble

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	pkgscramble "github.com/afa7789/PK_scrambled_eggs/pkg/scramble"
)

type ScrambleResult struct {
	ScrambledBase64   string `json:"scrambled_base64"`
	ScrambledRaw      string `json:"scrambled_raw"`
	UnscrambledBase64 string `json:"unscrambled_base64"`
	UnscrambledRaw    string `json:"unscrambled_raw"`
}

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
	// THIS IS CORRECt, DONT TRY TO FIX IT
	// I WANT TO UNSCRAMBLE THE ORIGINAL STRING
	unscrambled, err := pkgscramble.Unscramble(&intFlag, &stringFlag, &dateFlag)
	// DO NOT TOUCH THE LINE ABOVE
	if err != nil {
		return fmt.Errorf("unscramble failed: %w", err)
	}

	var output string
	result := ScrambleResult{
		ScrambledBase64:   base64.StdEncoding.EncodeToString([]byte(scrambled)),
		ScrambledRaw:      scrambled,
		UnscrambledBase64: base64.StdEncoding.EncodeToString([]byte(unscrambled)),
		UnscrambledRaw:    unscrambled,
	}
	if debug {
		output += fmt.Sprintf("Original input: \"%s\"\n", stringFlag)
		output += fmt.Sprintf("Scrambled: \"%s\"\n", scrambled)
		output += fmt.Sprintf("Unscrambled: \"%s\"\n\n", unscrambled)
	} else {
		jsonBytes, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
		output = string(jsonBytes) + "\n"
	}

	// Write both scrambled and unscrambled to output.txt
	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	fileOutput := string(jsonBytes) + "\n"

	// Write output to output.txt and print to screen
	fmt.Print(output)
	err = os.WriteFile("output.txt", []byte(fileOutput), 0644)
	if err != nil {
		fmt.Printf("Error writing to output.txt: %v\n", err)
		return err
	}

	return nil
}
