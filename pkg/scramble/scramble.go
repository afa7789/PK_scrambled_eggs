package scramble

import (
	"time"

	nconst "github.com/afa7789/PK_scrambled_eggs/pkg/natural_const"
)

// Scramble scrambles a string using the custom constant rules
func Scramble(intFlag *int, input *string, dateFlag *string) (string, error) {
	if intFlag == nil || input == nil || dateFlag == nil {
		panic("ScrambleString received nil pointer")
	}
	inputVal := *input
	dateVal := *dateFlag
	var date time.Time
	if dateVal == "00-00-0000" {
		date = time.Time{}
	} else {
		var err error
		date, err = time.Parse("01-02-2006", dateVal)
		if err != nil {
			return "", err
		}
	}
	var cc *nconst.CustomConstant
	if date.IsZero() {
		cc = nconst.NewCustomConstantDateZero(*intFlag)
	} else {
		cc = nconst.NewCustomConstant(*intFlag, date)
	}
	inputRunes := []rune(inputVal)
	cc.CalculateCustomConstant(len(inputRunes) + 20)
	scrambledArray := make([]rune, len(inputRunes))
	for i, char := range inputRunes {
		scrambledArray[i] = char + int32(cc.Digits[i])
	}
	return string(scrambledArray), nil
}

// Unscramble unscrambles a string using the custom constant rules
func Unscramble(intFlag *int, scrambled *string, dateFlag *string) (string, error) {
	if intFlag == nil || scrambled == nil || dateFlag == nil {
		panic("UnscrambleString received nil pointer")
	}
	scrambledVal := *scrambled
	dateVal := *dateFlag
	var date time.Time
	if dateVal == "00-00-0000" {
		date = time.Time{}
	} else {
		var err error
		date, err = time.Parse("01-02-2006", dateVal)
		if err != nil {
			return "", err
		}
	}
	var cc *nconst.CustomConstant
	if date.IsZero() {
		cc = nconst.NewCustomConstantDateZero(*intFlag)
	} else {
		cc = nconst.NewCustomConstant(*intFlag, date)
	}
	scrambledRunes := []rune(scrambledVal)
	cc.CalculateCustomConstant(len(scrambledRunes) + 20)
	unscrambledArray := make([]rune, len(scrambledRunes))
	for i, char := range scrambledRunes {
		unscrambledArray[i] = char - int32(cc.Digits[i])
	}
	return string(unscrambledArray), nil
}
