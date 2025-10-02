package scramble

import (
	"fmt"
	"os"
	"time"

	nconst "github.com/afa7789/PK_scrambled_eggs/pkg/natural_const"
)

// Adjust the import path as needed

func Scramble(
	intFlag int,
	stringFlag string,
	dateFlag string,
	debug bool, // default is false
) error {
	// Function body remains unchanged
	var date time.Time

	if dateFlag == "00-00-0000" {
		date = time.Time{}
	} else {
		var err error
		date, err = time.Parse("01-02-2006", dateFlag)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			os.Exit(1)
		}
	}
	var cc *nconst.CustomConstant
	if date.IsZero() {
		cc = nconst.NewCustomConstantDateZero(intFlag)
	} else {
		cc = nconst.NewCustomConstant(intFlag, date)
	}
	cc.CalculateCustomConstant(len(stringFlag) + 20)
	str := cc.String()
	if debug {
		fmt.Printf("Custom constant[%d]: %s\n\n", len(str), str)
	}
	scrambledArray := make([]rune, len(stringFlag))

	for i, char := range stringFlag {
		newRune := char + int32(cc.Constant.Digits[i])
		if debug {
			fmt.Printf("Scramble -> %c -> : %c , %03d + %01d = %d \n", char, newRune, char, int32(cc.Constant.Digits[i]), newRune)
		}
		scrambledArray[i] = char + int32(cc.Constant.Digits[i])
	}
	var output string
	if debug {
		output += fmt.Sprintf("Scrambled string \"%s\" to \"%s\"\n\n", stringFlag, string(scrambledArray))
	} else {
		output += fmt.Sprintf("Scrambled: \"%s\"\n\n", string(scrambledArray))
	}

	unscrambledArray := make([]rune, len(stringFlag))
	for i, char := range stringFlag {
		newRune := char - int32(cc.Constant.Digits[i])
		if debug {
			fmt.Printf("Unscramble -> %c -> : %c , %03d - %01d = %d \n", char, newRune, char, int32(cc.Constant.Digits[i]), newRune)
		}
		unscrambledArray[i] = char - int32(cc.Constant.Digits[i])
	}
	if debug {
		output += fmt.Sprintf("Unscrambled string \"%s\" to \"%s\"\n", string(stringFlag), string(unscrambledArray))
	} else {
		output += fmt.Sprintf("Unscrambled: \"%s\"\n", string(unscrambledArray))
	}
	// Write output to output.txt and print to screen
	fmt.Print(output)
	err := os.WriteFile("output.txt", []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error writing to output.txt: %v\n", err)
		return err
	}

	return nil
}
