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
) error {

	var date time.Time

	if dateFlag == "00-00-0000" {
		date = time.Time{}
	} else {
		var err error
		// Here you can add the logic to process the flags
		// For example, you can create a new custom constant based on the integer flag
		// and the date flag
		date, err = time.Parse("01-02-2006", dateFlag)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			os.Exit(1)
		}
	}
	var cc *nconst.CustomConstant
	if date.IsZero() {
		cc = nconst.NewCustomConstantDateZero(intFlag) // using a natural constant as the base ( on int flag )
	} else {
		// using a natural constant as the base ( on int flag )
		cc = nconst.NewCustomConstant(intFlag, date)
	}
	cc.CalculateCustomConstant(len(stringFlag) + 20) // entry is the number of digits to calculate
	str := cc.String()
	fmt.Printf("Custom constant[%d]: %s\n\n", len(str), str)

	// with the string flag, you can do something else, create array of same length
	// and scramble it with the digits of the custom constant
	newArray := make([]rune, len(stringFlag))
	// for each character in the string, scramble it with the digit of the custom constant in the same index
	for i, char := range stringFlag {
		newRune := char + int32(cc.Constant.Digits[i])
		fmt.Printf("Scramble -> %c -> : %c , %03d + %01d = %d \n", char, newRune, char, int32(cc.Constant.Digits[i]), newRune)
		newArray[i] = char + int32(cc.Constant.Digits[i])
	}

	fmt.Printf("Scrambled string \"%s\" to \"%s\"\n\n", stringFlag, string(newArray))

	newArray2 := make([]rune, len(stringFlag))
	// can we revert here to make sure it works ?
	for i, char := range stringFlag {
		newRune := char - int32(cc.Constant.Digits[i])
		fmt.Printf("Unscramble -> %c -> : %c , %03d - %01d = %d \n", char, newRune, char, int32(cc.Constant.Digits[i]), newRune)
		newArray2[i] = char - int32(cc.Constant.Digits[i])
	}
	fmt.Printf("Unscrambled string \"%s\" to \"%s\"\n", string(stringFlag), string(newArray2))

	return nil
}
